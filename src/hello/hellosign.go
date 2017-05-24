package main

import (
	"fmt"
	"strings"
	"net/http"
	"os"
	"io/ioutil"
	"bytes"
	"log"
	"mime/multipart"
)

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params string, paramName, path string) (*http.Request, error) {
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
	url := "https://50cff1d3451e96e91333185a91f6a5e1ccab34094fe2424a1ecbe88e580192ae:@api.hellosign.com/v3/signature_request/send"

	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"message\"\r\n\r\nmessage here\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"title\"\r\n\r\ntitle here\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"subject\"\r\n\r\nsubject here\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"file[0]\"; filename=\"11.docx\"\r\nContent-Type: application/vnd.openxmlformats-officedocument.wordprocessingml.document\r\n\r\n\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"signers[0][email_address]\"\r\n\r\nalex+signer@hellosign.com\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"signers[0][name]\"\r\n\r\nalex\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")
	req, err := http.newfileUploadRequest(url, payload, "file", "/10.docx")

	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(res)
		fmt.Println(string(body))
	}

}
