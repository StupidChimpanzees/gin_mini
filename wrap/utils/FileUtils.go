package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/http"
	"os"
)

func GetFileType(file any) (string, error) {
	buffer := make([]byte, 512)
	if value, ok := file.(string); ok {
		fileStream, err := os.Open(value)
		if err != nil {
			return "", err
		}
		defer func(fileStream *os.File) {
			err := fileStream.Close()
			if err != nil {
				return
			}
		}(fileStream)
		_, err = fileStream.Read(buffer)
		if err != nil {
			return "", err
		}
	} else if value, ok := file.([]byte); ok {
		buffer = value[:512]
	} else {
		return "", errors.New("parameter type error")
	}

	return http.DetectContentType(buffer), nil
}

func GetSmallFileContent(file string) ([]byte, error) {
	fileStream, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func(fileStream *os.File) {
		err = fileStream.Close()
		if err != nil {
			return
		}
	}(fileStream)
	return io.ReadAll(fileStream)
}

func GetRequestParams(c *gin.Context) *[]byte {
	var b []byte
	if c.Request.Method == "GET" || c.Request.Method == "DELETE" {
		b, _ = json.Marshal(c.Request.URL.Query())
	} else {
		switch c.ContentType() {
		case binding.MIMEJSON, binding.MIMEXML, binding.MIMEXML2,
			binding.MIMEPROTOBUF, binding.MIMEMSGPACK, binding.MIMEMSGPACK2:
			b, _ = c.GetRawData()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
		case binding.MIMEMultipartPOSTForm:
			b, _ = json.Marshal(c.Request.PostForm)
		default:
			b, _ = json.Marshal(c.Request.PostForm)
		}
	}
	return &b
}
