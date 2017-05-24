package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	// "encoding/json"
)

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/sigreqwtemplate", handler2)

	http.ListenAndServe(":6789", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>BUTTONS THAT DO THINGS</h1><br /><p><a href=\"/sigreqwtemplate\">Click Here</a> to trigger a templated signature request.</p>")

}

func handler2(w http.ResponseWriter, r *http.Request) {

	url := "https://:@api.hellosign.com/v3/signature_request/send_with_template"

	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"client_id\"\r\n\r\nd7219512693825facdd9241f458decf2\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"template_id\"\r\n\r\n06874015c4197f5b8d7a03f159095385d89c74d5\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"test_mode\"\r\n\r\n1\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"subject\"\r\n\r\nusing postman\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"message\"\r\n\r\ntemplate non embedded\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"signers[Signer][email_address]\"\r\n\r\nalex+postman1@hellosign.com\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"signers[Signer][name]\"\r\n\r\nMr Postman\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"custom_fields\"\r\n\r\n[{ \"name\":\"label1\", \"value\":\"things here aren't always what they seem\" } ]\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")

	req, err := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if err != nil {
		//TODO figure out how to get the HTTP response code from the call, because the server's returning a 401 but my code still thinks
		//the call was successful
		fmt.Println(err)
		fmt.Fprintf(w, "<p>Crud - there was an error. Server logs will show it.</p><br />")
		fmt.Println(res)
		fmt.Println(string(body))
		fmt.Fprintf(w, "<br /><p><a href=\"/\">Click Here</a> to go home.</p>")
	} else {
	fmt.Fprintf(w, "<p>It was successfull!</p><br />")
	fmt.Println(res)
	// apiresponse := map[string]int{body}
	// jsonresponse, _ := json.Marshal(apiresponse)
	// fmt.Println(jsonresponse)
	fmt.Println(string(body))
	fmt.Fprintf(w, "<br /><p><a href=\"/\">Click Here</a> to go home.</p>")
	}
}
