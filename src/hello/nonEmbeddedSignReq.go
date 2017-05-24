package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", uri, body)
}

func main() {
	path, _ := os.Getwd()
	path += "/10.docx"
	extraParams := map[string]string{
		"title":       "My Document",
		"message":      "message here",
		"signer[0][name]": "alex",
		"signer[0][email_address]": "alex@hellosign.com",
		"test_mode": "true",
	}
	request, err := newfileUploadRequest("https://50cff1d3451e96e91333185a91f6a5e1ccab34094fe2424a1ecbe88e580192ae:@api.hellosign.com/v3/signature_request/send", extraParams, "file", "10.docx")
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(resp)
		fmt.Println(string(body))
	}
}
