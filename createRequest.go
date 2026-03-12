package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func createRequest(caller HCACaller, mainHCA HCAMain) (RequestReturn, error) {
	url := mainHCA.BaseURL + caller.Path

	var req *http.Request
	if caller.Method == "GET" {
		req, _ = http.NewRequest("GET", url, nil)
	} else {
		jsonData, _ := json.Marshal(caller.JSON)
		req, _ = http.NewRequest(caller.Method, url, strings.NewReader(string(jsonData)))
		req.Header.Set("Content-Type", "application/json")
	}

	for k, v := range mainHCA.Headers {
		req.Header.Set(k, v)
	}
	for k, v := range caller.Headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return RequestReturn{}, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return RequestReturn{
		Body:   body,
		Status: resp.Status,
	}, nil
}
