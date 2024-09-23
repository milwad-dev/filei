package filei

import (
	"io"
	"mime/multipart"
	"os"
)

// UploadFile saves an uploaded file to the specified path.
func UploadFile(file multipart.File, fileName string) error {
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}
