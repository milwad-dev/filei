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

func TestExistsFile(t *testing.T) {
	isExists := Exists("testdata/text.txt") // Must be true

	if isExists != true {
		t.Error("Exists() error data must be true")
	}

	isNotExists := Exists("wrong path") // Must be false
	if isNotExists != false {
		t.Error("Exists() error data must be false")
	}
}

func TestCleanDirectory(t *testing.T) {
	directory := "testdata/dir"

	defer os.Create(directory + "/data.txt")
	defer os.Create(directory + "/item.txt")

	dirData, _ := os.ReadDir(directory)

	if len(dirData) != 2 {
		t.Fatal("CleanDirectory() error data must have 2 files")
	}

	CleanDirectory(directory)

	dirData, _ = os.ReadDir(directory)

	if len(dirData) == 2 {
		t.Fatal("CleanDirectory() error data must have 0 files")
	}
}

func TestDeleteDirectory(t *testing.T) {
	directory := "testdata/dir2"

	defer os.Mkdir(directory, 0755)

	err := DeleteDirectory(directory)

	if err != nil {
		t.Fatalf("DeleteDirectory() error = %v", err)
	}
}

func TestMoveFile(t *testing.T) {
	sourcePath := "testdata/move-text.txt"
	destPath := "testdata/move2-text.txt"

	func() {
		err := MoveFile(sourcePath, destPath)
		if err != nil {
			t.Fatalf("MoveFile() error = %v", err)
		}
	}()

	defer MoveFile(destPath, sourcePath)

	// Wrong path
	err := MoveFile("wrong path", "wrong path")
	if err == nil {
		t.Fatalf("MoveFile() error data must have a value")
	}
}

func TestFiles(t *testing.T) {
	directory := "testdata"

	data, err := Files(directory)
	if err != nil {
		t.Fatalf("Files() error = %v", err)
	}

	if len(data) != 7 {
		t.Fatal("Files() error data must have 7 files")
	}
}

func TestSize(t *testing.T) {
	filePath := "testdata/text.txt"

	size, err := Size(filePath)
	if err != nil {
		t.Fatalf("Size() error = %v", err)
	}

	if size != 13 {
		t.Fatal("Size() error data must have 13 files")
	}
}

func TestChmod(t *testing.T) {
	filePath := "testdata/text.txt"

	err := Chmod(filePath, 0700)
	if err != nil {
		t.Errorf("Chmod() error = %v", err)
	}

	err = Chmod("wrong path", 0700)
	if err == nil {
		t.Errorf("Chmod() error data must have a value")
	}
}
