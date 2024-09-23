package filei

import (
	"os"
	"testing"
)

func TestUploadFile(t *testing.T) {
	file, _ := os.Create("testdata/new-text.txt")

	err := UploadFile(file, "testdata/upload.txt")
	if err != nil {
		t.Errorf("UploadFile() error = %v", err)
	}

	var fakeFile *os.File

	err = UploadFile(fakeFile, "testdata/upload.txt")
	if err == nil {
		t.Errorf("UploadFile() error data must have a value")
	}
}

func TestGetFile(t *testing.T) {
	_, err := GetFile("testdata/text.txt")
	if err != nil {
		t.Errorf("GetFile() error = %v", err)
	}

	_, err = GetFile("wrong path")
	if err == nil {
		t.Errorf("GetFile() error data must have a value")
	}
}

func TestDeleteFile(t *testing.T) {
	err := DeleteFile("testdata/delete-text.txt")
	if err != nil {
		t.Errorf("DeleteFile() error = %v", err)
	}

	err = DeleteFile("wrong path")
	if err == nil {
		t.Errorf("DeleteFile() error data must have a value")
	}

	os.Create("testdata/delete-text.txt")
}

func TestCreateFile(t *testing.T) {
	data := []byte("milwad")

	err := CreateFile("testdata/new-text.txt", data)
	if err != nil {
		t.Errorf("CreateFile() error = %v", err)
	}
}
