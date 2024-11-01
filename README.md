# FileI

The `FileI` is a lightweight and easy-to-use package for handling file operations in Go (Golang). It simplifies common file-related tasks such as uploading, deleting, and managing files, allowing developers to handle file I/O operations with minimal effort. Whether you're building an API, managing assets, or simply need a streamlined way to work with files in Go, FileI provides intuitive methods that integrate seamlessly into your project.

[![GoDoc Widget](https://godoc.org/github.com/milwad-devi/filei?status.svg)](https://pkg.go.dev/github.com/milwad-dev/filei)
[![Passed Tests](https://github.com/milwad-dev/filei/actions/workflows/tests.yml/badge.svg)](https://github.com/milwad-dev/filei/actions/workflows/tests.yml)

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
- [Prepend](#prepend)
- [Append](#append)

<a name="upload-file"></a>
### UploadFile

If you want to upload a file to a path you can use `UploadFile` function:

```go
file, _ := os.Create("testdata/new-text.txt")

err := UploadFile(file, "destPath")
if err != nil {
    // ...
}
```

<a name="get-file"></a>
### GetFile

If you want to retrieve a file from the specified path, you can use `GetFile` function:

```go
file, err := GetFile("file path")
if err != nil {
	// ...
}
```

<a name="delete-file"></a>
### DeleteFile

If you want to delete a file, you can use `DeleteFile` function:

```go
err := DeleteFile("file path")
if err != nil {
    // ...
}
```

<a name="create-file"></a>
### CreateFile

If you want to create a file with specific extension and contains a data, you can use `CreateFile` function:

```go
data := []byte("milwad")

err := CreateFile("file path", data)
if err != nil {
    // ...
}
```

<a name="exists"></a>
### Exists

If you want to ensure that a file exists, you can use `Exists` function:

```go
isExists := Exists("file path") // Bool
```

<a name="clean-directory"></a>
### CleanDirectory

If you want to clean a directory (remove all files in a directory), you can use `CleanDirectory` function:

```go
CleanDirectory("directory path") // Bool
```

<a name="delete-directory"></a>
### DeleteDirectory

If you want to delete a directory, you can use `DeleteDirectory` function:

```go
err := DeleteDirectory(directory)
```

<a name="move-directory"></a>
### MoveDirectory

If you want to move a directory to another path, you can use `MoveDirectory` function:

```go
err := MoveFile(sourcePath, destPath)
```

<a name="files"></a>
### Files

If you want to get all files of a directory, you can use `Files` function:

```go
files, err := Files(directory)
```

<a name="size"></a>
### Size

If you want to get a size of a file, you can use `Size` function:

```go
size, err := Size(filePath)
```

<a name="chmod"></a>
### Chmod

If you want to give permission to a file, you can use `Chmod` function:

```go
err := Chmod(filePath, 0700)
```

<a name="prepend"></a>
### Prepend

If you want to write to the beginning of a file, you can use `Prepend` function:

```go
err := Prepend(filePath, data)
```

<a name="append"></a>
# Append

If you want to write to the end of a file, you can use `Append` function:

```go
err := Append(filePath, data)
```

## License

* This package is created and modified by <a href="https://github.com/milwad-dev" target="_blank">Milwad Khosravi</a> for Laravel upper more than 9 and has been released under the MIT License.

## Contributing

This project exists thanks to all the people who contribute. [CONTRIBUTING](https://github.com/milwad-dev/filei/graphs/contributors)

<a href="https://github.com/milwad-dev/filei/graphs/contributors"><img src="https://opencollective.com/filei/contributors.svg?width=890&button=false" /></a>

## Security

If you've found a bug regarding security please mail [milwad.dev@gmail.com](mailto:milwad.dev@gmail.com) instead of using the issue tracker.


## Donate

If this package is helpful for you, you can buy a coffee for me :) ❤️

- Iraninan Gateway: https://daramet.com/milwad_khosravi
- Paypal Gateway: SOON
- MetaMask Address: `0xf208a562c5a93DEf8450b656c3dbc1d0a53BDE58`
