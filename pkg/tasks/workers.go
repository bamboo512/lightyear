package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"lightyear/core/global"
	"lightyear/models"
	"lightyear/pkg/transcoder"
	"log"
	"path/filepath"
	"time"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeImageTranscode = "image:transcode"
)

type ImageTranscodePayload struct {
	UUID     string
	encoding string
	quality  int
}

//----------------------------------------------
// Write a function NewXXXTask to create a task.
// A task consists of a type and a payload.
//----------------------------------------------

func NewImageTranscodeTask(uuid string, encoding string, quality int) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageTranscodePayload{UUID: uuid, encoding: encoding, quality: quality})
	if err != nil {
		return nil, err
	}
	// task options can be passed to NewTask, which can be overridden at enqueue time.
	return asynq.NewTask(TypeImageTranscode, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
}

//---------------------------------------------------------------
// Write a function HandleXXXTask to handle the input task.
// Note that it satisfies the asynq.HandlerFunc interface.
//

// Handler doesn't need to be a function. You can define a type
// that satisfies asynq.Handler interface. See examples below.
//---------------------------------------------------------------

// ImageProcessor implements asynq.Handler interface.
type ImageProcessor struct {
	// ... fields for struct
}

func (processor *ImageProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p ImageTranscodePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Transcoding image: uuid=%s", p.UUID)

	// Image transcode ...
	file, err := models.GetFileByUUID(p.UUID)
	if err != nil {
		return fmt.Errorf("error: failed to get file: %v", err)
	}
	// Update file instance
	file.ExpectedExtension = "." + p.encoding

	// Transcode the image and handle any errors
	originalPath := filepath.Join(global.Config.Storage.OriginalImagePath, p.UUID+file.OriginalExtension)
	encodedPath := filepath.Join(global.Config.Storage.TranscodedImagePath, p.UUID+file.ExpectedExtension)

	transcoder.TranscodeImage(originalPath, encodedPath, p.encoding, p.quality)

	if err != nil {
		log.Printf("Failed to transcode image: %v", err)
		return fmt.Errorf("error: failed to transcode image: %v", err)
	}

	file.FileStatus = models.FileConverted

	// Update File Status in the database
	err = models.UpdateFile(&file)
	if err != nil {
		return fmt.Errorf("error: failed to save to db")
	}

	log.Printf("Image transcoded successfully: uuid=%s", p.UUID)
	return nil
}

func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{}
}
