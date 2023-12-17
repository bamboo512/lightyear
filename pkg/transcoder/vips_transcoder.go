package transcoder

import (
	"errors"
	"fmt"
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

// TranscodeImage reads an image from the originalPath, encodes it into the specified encoding,
// and writes the encoded image to encodedPath.
func TranscodeImage(originalPath, encodedPath, encoding string, quality int) error {
	log("Reading image: ", originalPath)
	img, err := vips.NewImageFromFile(originalPath)
	if err != nil {
		return fmt.Errorf("failed to read image: %w", err)
	}
	defer img.Close()

	out, err := encodeImage(encoding, img, quality)
	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	if err := os.WriteFile(encodedPath, out, 0644); err != nil {
		return fmt.Errorf("failed to write encoded image: %w", err)
	}

	log("Image transcoded successfully!")
	return nil
}

// EncodeImage encodes an image into the specified format with a given quality.
func encodeImage(coding string, img *vips.ImageRef, quality int) (out []byte, err error) {
	if quality <= 0 || quality > 100 {
		return nil, errors.New("invalid quality")
	}

	switch coding {
	case "avif":
		exportParam := vips.AvifExportParams{Quality: quality}
		out, _, err = img.ExportAvif(&exportParam)

	case "heic":
		exportParam := vips.HeifExportParams{Quality: quality}
		out, _, err = img.ExportHeif(&exportParam)

	default:
		exportParam := vips.HeifExportParams{Quality: quality}
		out, _, err = img.ExportHeif(&exportParam)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to encode image: %w", err)
	}

	return out, nil
}

// log is a simple logging function that could be further expanded or replaced with a proper logging library.
func log(v ...any) {
	fmt.Println(v...)
}
