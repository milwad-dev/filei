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

// GetFile retrieves a file from the specified path.
func GetFile(fileName string) (multipart.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// DeleteFile removes a file from the specified path.
func DeleteFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}

// CreateFile create a new file from the specified path and data.
func CreateFile(filaName string, data []byte) error {
	file, err := os.Create(filaName)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
