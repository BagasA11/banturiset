package helpers

import (
	"errors"
	"net/http"
)

const oneMB = 1024 * 1024

// const pdf = "pdf"
// const jpg = "jpg"
// const jpeg = "jpeg"
// const png = "png"

func ValidateFileSize(r *http.Request) (bool, error) {

	r.ParseMultipartForm(10 << 20) // Set maximum file size to 10 MB

	_, fileheader, err := r.FormFile("file")
	if err != nil {
		return false, err
	}

	if fileheader.Size > 5*oneMB {
		return false, errors.New("file melebihi 1 mb")
	}

	return true, nil
}
