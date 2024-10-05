package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
)

func UploadFileFromUrl(client *http.Client, url string, fileName string) error {
	// Get the data
	gifResp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer gifResp.Body.Close()

  body := new(bytes.Buffer) // for the multipart data request
  multipartWriter := multipart.NewWriter(body)

  // The standard multipart file writer sets content type to 'application/octet' and does not set transfer encoding, so 
  // we write our own Part data.
  // part, err := multipartWriter.CreateFormFile("file", fileName)
  h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", FileContentDisposition("file", fileName))
	h.Set("Content-Type", "image/gif")
	h.Set("Content-Transfer-Encoding", "binary")
	part, err := multipartWriter.CreatePart(h)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, gifResp.Body)
	if err != nil {
		return err
	}

	multipartWriter.Close()

	uploadReq, err := http.NewRequest("POST", "http://backend:8082/api/v1/files/upload/", body)
	uploadReq.Header.Add("Content-Type", multipartWriter.FormDataContentType())

  uploadResponse, err := client.Do(uploadReq)
  // defer uploadResponse.Body.Close()
  if err != nil {
  	return err
  }

  statusCode := uploadResponse.StatusCode
  if statusCode > 399 {
  	buf := new(strings.Builder)
		io.Copy(buf, uploadResponse.Body)
		message := fmt.Sprintln(buf.String())
  	return fmt.Errorf("Upload attempt failed; %d: %s", statusCode, message)
  }

	return nil
}

func FileContentDisposition(fieldname, filename string) string {
    return fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldname, filename)
}