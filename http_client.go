package main

import (
	"net/http"
	"net/url"
	"encoding/json"
	"log"
)

func get(client *http.Client, url string, values url.Values, typedef interface{}) interface{} {

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Client-ID", "<X-Client-ID>")
	req.Header.Add("X-Access-Token", "<X-Access-Token>")
	req.URL.RawQuery = values.Encode()
	resp, _ := client.Do(req)

	error := json.NewDecoder(resp.Body).Decode(&typedef)
	if error != nil {
		log.Fatal(error)
	}

	return typedef
}

