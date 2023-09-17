package weather

import (
	"io"
	"net/http"
)

const JmaApiEndpoint = "https://www.jma.go.jp/bosai/forecast/data/forecast/140000.json"

// FetchDataFromJma makes a GET request to the JMA API, reads the response body, and returns it.
func FetchDataFromJma() ([]byte, error) {
	resp, err := http.Get(JmaApiEndpoint)
	if err != nil {
		return nil, err
	}
	// Close the response body once all operations on it are done.
	// This is essential to release resources and avoid potential memory leaks.
	defer resp.Body.Close()

	// Read the entire response body and return its contents.
	return io.ReadAll(resp.Body)
}
