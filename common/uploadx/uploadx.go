package uploadx

import (
	"io"
	"mime/multipart"
	"os"
	"path"
)

//
//func UploadFile(r *http.Request, maxFileSize int64, headerKeyName, filePath string) (string, error) {
//	err := r.ParseMultipartForm(maxFileSize)
//	if err != nil {
//		return "", err
//	}
//
//	file, handler, err := r.FormFile(headerKeyName)
//	if err != nil {
//		return "", err
//	}
//
//	defer file.Close()
//
//	logx.Infof("Upload file: %+v, file size: %v, header: %+v",
//		handler.Filename, handler.Size, handler.Header,
//	)
//
//	temp, err := os.Create(path.Join(filePath, handler.Filename))
//	if err != nil {
//		return "", err
//	}
//
//	defer temp.Close()
//
//	io.Copy(temp, file)
//	return handler.Filename, nil
//}

func UploadFile(f multipart.File, header *multipart.FileHeader, filePath string) (string, error) {
	temp, err := os.Create(path.Join(filePath, header.Filename))
	if err != nil {
		return "", err
	}

	defer temp.Close()

	_, _ = io.Copy(temp, f)

	return header.Filename, nil
}
