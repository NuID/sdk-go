package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

// Each call to the API requires an X-API-Key header, and the correct
// content negotiation headers for JSON.
func (auth *APIClient) defaultHeaders(req *http.Request) {
	req.Header.Add("X-API-Key", auth.ApiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}

// Perform a GET request against the API at the given path.
func (auth *APIClient) get(path string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", auth.path(path), nil)
	if err != nil {
		return nil, err
	}

	auth.defaultHeaders(req)
	return http.DefaultClient.Do(req)
}

// Concat the api host and path together.
func (auth *APIClient) path(path string) string {
	return fmt.Sprintf("%s%s", auth.Host, path)
}

// Perform a POST against the API at the given path with the map converted to
// JSON.
func (auth *APIClient) post(path string, data map[string]interface{}) (resp *http.Response, err error) {
	body, err := toJSON(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", auth.path(path), body)
	if err != nil {
		return nil, err
	}

	auth.defaultHeaders(req)
	return http.DefaultClient.Do(req)
}

func toJSON(data map[string]interface{}) (io.Reader, error) {
	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return io.Reader(bytes.NewReader(buf)), nil
}
