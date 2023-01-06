package uploadx

import (
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"os"
	"testing"
)

func Test_Update(t *testing.T) {
	testCases := []struct {
		Name     string
		filePath string
		testOpt  func(t *testing.T, path string)
	}{
		{
			Name:     "Success",
			filePath: "../../resources",
			testOpt: func(t *testing.T, path string) {
				file, err := os.Open("../../resources/test/testing.txt")
				assert.Nil(t, err)
				defer file.Close()

				info, err := file.Stat()
				assert.Nil(t, err)

				header := &multipart.FileHeader{
					Filename: info.Name(),
					Size:     info.Size(),
				}

				fileName, err := UploadFile(file, header, path)
				assert.Nil(t, err)
				assert.NotEmpty(t, fileName)
			},
		},
		{
			Name:     "path exist",
			filePath: "../../resources",
			testOpt: func(t *testing.T, path string) {
				_, err := os.Open("../resources/test/testing.txt")
				assert.NotNil(t, err)
			},
		},
		{
			Name:     "file exist",
			filePath: "../../resources",
			testOpt: func(t *testing.T, path string) {
				_, err := os.Open("../../resources/test/abc.txt")
				assert.NotNil(t, err)
			},
		},
		{
			Name:     "cannot find the path specified.",
			filePath: "../../test",
			testOpt: func(t *testing.T, path string) {
				file, err := os.Open("../../resources/test/testing.txt")
				assert.Nil(t, err)
				defer file.Close()

				info, err := file.Stat()
				assert.Nil(t, err)

				header := &multipart.FileHeader{
					Filename: info.Name(),
					Size:     info.Size(),
				}
				_, err = UploadFile(file, header, path)
				assert.NotNil(t, err)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			test.testOpt(t, test.filePath)
		})
	}
}
