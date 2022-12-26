package api

import (
	"io"
	"log"
	"net/http"
)

func request(source string) []byte {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", source, nil)
	if err != nil {
		log.Fatalln(err)
	}

	response, err := httpClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}
