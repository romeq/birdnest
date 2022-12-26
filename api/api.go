package api

import (
	"io"
	"net/http"
)

func request(source string) (body []byte, err error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", source, nil)
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
