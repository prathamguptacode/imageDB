package services

import (
	"fmt"
	"mime/multipart"
	"net/http"
)

func ValidateFileType(file *multipart.FileHeader, allowed []string) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	// Read the first 512 bytes for MIME detection
	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil {
		return err
	}

	mimeType := http.DetectContentType(buf[:n])

	for _, a := range allowed {
		if mimeType == a {
			return nil
		}
	}

	return fmt.Errorf("file type %s is not allowed", mimeType)
}
