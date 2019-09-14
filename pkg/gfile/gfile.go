package gfile

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"time"
)

type FileInfo = struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
	Sys     interface{}
	Md5     string
}

type FileSlice = struct {
	Base64Content string
	PartSize      int64
	PartMd5       string
}

type FileBase64 = struct {
	FileSize      int64
	Base64Content string
}

func GetFileInfo(filePath string) (fileInfo FileInfo, err error) {

	fileStat, err := os.Stat(filePath)

	if err != nil {
		return fileInfo, err
	}

	fileMd5, err := GetFileMd5(filePath)

	if err != nil {
		return fileInfo, err
	}

	fileInfo.Size = fileStat.Size()
	fileInfo.Name = fileStat.Name()
	fileInfo.IsDir = fileStat.IsDir()
	fileInfo.Mode = fileStat.Mode()
	fileInfo.ModTime = fileStat.ModTime()
	fileInfo.Sys = fileStat.Sys()

	fileInfo.Md5 = fileMd5

	return fileInfo, nil
}

func GetFileBase64Content(filePath string) (fileBase64 FileBase64, err error) {

	file, err := os.Open(filePath)

	if err != nil {
		return fileBase64, err
	}

	fileStat, err := file.Stat()

	if err != nil {
		return fileBase64, err
	}

	defer file.Close()

	fileBytes := make([]byte, fileStat.Size())

	n, err := file.Read(fileBytes)

	if err != nil {
		return fileBase64, err
	}

	fileBase64.FileSize = int64(n)
	fileBase64.Base64Content = base64.StdEncoding.EncodeToString(fileBytes)

	return fileBase64, nil
}

func GetFileMd5(filePath string) (fileMd5 string, err error) {

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	md5Obj := md5.New()

	_, err = io.Copy(md5Obj, file)

	if err != nil {
		return "", err
	}

	fileMd5 = hex.EncodeToString(md5Obj.Sum(nil))

	return fileMd5, nil
}

func GetSliceFile(filePath string, offset int64, size int64) (fileSlice FileSlice, err error) {

	file, err := os.Open(filePath)
	if err != nil {
		return fileSlice, err
	}

	defer file.Close()

	fileBytes := make([]byte, size)

	n, err := file.ReadAt(fileBytes, offset)

	if err != nil {
		if err == io.EOF {
			fmt.Println("here")
		}
		return fileSlice, err
	}

	fileSlice.Base64Content = base64.StdEncoding.EncodeToString(fileBytes)
	fileSlice.PartSize = int64(n)
	fileSlice.PartMd5 = GetBytesMd5(fileBytes)

	return fileSlice, nil
}

func GetBytesMd5(b []byte) (stringMd5 string) {

	md5Obj := md5.New()

	md5Obj.Write(b)

	stringMd5 = hex.EncodeToString(md5Obj.Sum(nil))

	return stringMd5
}
