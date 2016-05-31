package ehop

import (
	"bytes"
	"crypto/tls"
	"log"
	"net/http"
)

// CreateEhopRequest creates and sends HTTP request to ExtraHop system.  Returns the response
func CreateEhopRequest(method string, call string, payload string, APIKey string, path string) *http.Response {
	//Create a 'transport' object... this is necessary if we want to ignore
	//the EH insecure CA.  Similar to '--insecure' option for curl
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//Crate a new client object... and pass him the parameters of the transport object
	//we created above
	client := http.Client{Transport: tr}
	postBody := []byte(payload)
	req, err := http.NewRequest(method, path+call, bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalf("Failed to create HTTP request: %q", err.Error())
	}

	//Add some header stuff to make it EH friendly
	req.Header.Add("Authorization", APIKey)
	req.Header.Add("Content-Type", " application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to perform HTTP request: %q", err.Error())

	}
	return resp
}
