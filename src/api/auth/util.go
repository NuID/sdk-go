package api_auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

func (auth *AuthAPI) defaultHeaders(req *http.Request) {
	req.Header.Add("X-API-Key", auth.ApiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}

func (auth *AuthAPI) get(path string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", auth.path(path), nil)
	if err != nil {
		return nil, err
	}

	auth.defaultHeaders(req)
	return http.DefaultClient.Do(req)
}

func (auth *AuthAPI) path(path string) string {
	return fmt.Sprintf("%s%s", auth.Host, path)
}

func (auth *AuthAPI) post(path string, data map[string]interface{}) (resp *http.Response, err error) {
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
