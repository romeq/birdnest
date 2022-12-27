package api

import (
	"io"
	"net/http"
)

func request(source string) (body []byte, err error) {
	resp, err := http.Get(source)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
