package info

import (
	"encoding/json"
	"github.com/katy248/gravatar/pkg/url"
	"io"
	"log"
	"net/http"
)

func parseResponse(response []byte) Response {
	var result Response
	err := json.Unmarshal(response, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func FetchData(email string) *ProfileInfo {
	link := url.NewInfoUrl(email, url.InfoFormat(url.JsonFormat))

	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	r := parseResponse(body)
	info := r.GetInfo()

	return &info
}
