package ffmpeg

import (
	"fmt"
	"os/exec"
)

func TranscodeImage(originalPath, encodedPath, encoding string, quality int) error {
	exec.Command("ffmpeg", "-i", originalPath, "-q:v", fmt.Sprintf("%d", quality), encodedPath)
	return nil
}
