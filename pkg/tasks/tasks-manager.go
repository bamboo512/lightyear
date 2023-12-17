package tasks

import (
	"lightyear/core/global"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func AddNewTask(uuid string, encoding string, quality int) {
	task, err := NewImageTranscodeTask(uuid, encoding, quality)
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	// Enqueue the task with options.
	info, err := global.AsynqClient.Enqueue(task,
		// asynq.ProcessIn(24*time.Hour),   // process in the future
		asynq.MaxRetry(10),
		asynq.Timeout(3*time.Minute),
	)

	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
