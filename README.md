# FileI

The `FileI` is a lightweight and easy-to-use package for handling file operations in Go (Golang). It simplifies common file-related tasks such as uploading, deleting, and managing files, allowing developers to handle file I/O operations with minimal effort. Whether you're building an API, managing assets, or simply need a streamlined way to work with files in Go, FileI provides intuitive methods that integrate seamlessly into your project.

[![GoDoc Widget](https://godoc.org/github.com/milwad-devi/filei?status.svg)](https://pkg.go.dev/github.com/milwad-dev/filei)

Key Features:

- File Upload: Easily upload files with minimal configuration.
- File Deletion: Remove files from your system with a single function call.
- Error Handling: Provides built-in error handling for common file operations.
- Cross-platform Support: Works across different operating systems.

Perfect for developers who need to manage file operations quickly and efficiently in their Go projects.

## Installation

For install the `FileI` on your project, you can run below command on terminal:

```shell
go get -u github.com/milwad-dev/filei
```

## Usage

- [UploadFile](#upload-file)
- [GetFile](#get-file)
- [DeleteFile](#delete-file)
- [CreateFile](#create-file)
- [Exists](#exists)
- [CleanDirectory](#clean-directory)
- [DeleteDirectory](#delete-directory)
- [MoveFile](#move-file)
- [Files](#files)
- [Size](#size)
- [Chmod](#chmod)

<a name="upload-file"></a>
### UploadFile

If you want to upload a file to a path you can use `UploadFile` function:

```go
file, _ := os.Create("testdata/new-text.txt")

err := UploadFile(file, "testdata/upload.txt")
if err != nil {
    // ...
}
```
