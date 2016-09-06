package web

import (
	"log"
	"net/http"
)

//SendRequest Used To Send HTTP request to website
func SendRequest(url string) int {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return resp.StatusCode
}
