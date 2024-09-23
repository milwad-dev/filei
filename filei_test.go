package filei

import "testing"

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

func TestCreateFile(t *testing.T) {
	data := []byte("milwad")

	err := CreateFile("testdata/new-text.txt", data)
	if err != nil {
		t.Errorf("CreateFile() error = %v", err)
	}
}
