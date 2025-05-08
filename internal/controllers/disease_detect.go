package controllers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DiseaseDetect(c *gin.Context) {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return
	}

	defer file.Close()

	body := &bytes.Buffer{}
	write := multipart.NewWriter(body)

	part, err := write.CreateFormFile("image", header.Filename)
	if err != nil {
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return
	}

	write.Close()

	url := "http://192.168.1.12:8000/detect-plant-disease"
	method := "POST"

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", write.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), responseBody)

}
