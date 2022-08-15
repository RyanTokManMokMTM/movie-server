package uploadx

import (
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"os"
	"path"
)

func UploadFile(r *http.Request, maxFileSize int64, headerKeyName, filePath string) (string, error) {
	r.ParseMultipartForm(maxFileSize)
	file, handler, err := r.FormFile(headerKeyName)
	if err != nil {
		return "", err
	}

	defer file.Close()

	logx.Info("Upload file: %+v, file size: %v, header: %+v",
		handler.Filename, handler.Size, handler.Header,
	)

	temp, err := os.Create(path.Join(filePath, handler.Filename))
	if err != nil {
		return "", err
	}

	defer temp.Close()

	io.Copy(temp, file)
	return handler.Filename, nil
}
