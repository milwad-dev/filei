package filei

import (
	"fmt"
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

// GetFile retrieve a file from the specified path.
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

// Exists ensure that the file exists.
func Exists(fileName string) bool {
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

// CleanDirectory clear the directory.
func CleanDirectory(directory string) bool {
	data, _ := os.ReadDir(directory)

	for _, file := range data {
		err := os.Remove(directory + "/" + file.Name())
		if err != nil {
			continue
		}
	}

	return true
}

// DeleteDirectory delete the directory.
func DeleteDirectory(directory string) error {
	err := os.RemoveAll(directory)
	if err != nil {
		return err
	}

	return nil
}

// MoveFile move the file from one place to another place.
func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("couldn't open dest file: %v", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("couldn't copy to dest from source: %v", err)
	}

	inputFile.Close()

	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't remove source file: %v", err)
	}
	return nil
}

// Files get the list of files of a directory.
func Files(directoryPath string) ([]os.DirEntry, error) {
	data, err := os.ReadDir(directoryPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't read directory: %v", err)
	}

	return data, nil
}

// Size get the size of a file.
func Size(filePath string) (int64, error) {
	data, err := os.Stat(filePath)
	if err != nil {
		return -1, fmt.Errorf("couldn't stat file: %v", err)
	}

	return data.Size(), nil
}

// Chmod give permission to a file.
func Chmod(filePath string, mode os.FileMode) error {
	err := os.Chmod(filePath, mode)
	if err != nil {
		return fmt.Errorf("couldn't chmod file: %v", err)
	}

	return nil
}
