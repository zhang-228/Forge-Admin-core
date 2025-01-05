package file

import (
	"ForgeAdmin/pkg/xerr"
	"mime/multipart"
	"net/http"
)

// File 文件结构体
type File struct {
	File     multipart.File
	FileName string
	Size     int64
}

// FormFile 从请求中获取文件
func FormFile(r *http.Request, key string, maxSize int64) (*File, error) {
	file, header, err := r.FormFile(key)
	if err != nil {
		return nil, err
	}

	if header.Size/1024 > maxSize {
		return nil, xerr.UploadFileMaxSize
	}

	return &File{
		File:     file,
		FileName: header.Filename,
		Size:     header.Size,
	}, err
}
