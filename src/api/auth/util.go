package api_auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

var (
	client *http.Client
)

func Init() {
	client = &http.Client{}
}

func (auth *AuthAPI) defaultHeaders(req *http.Request) {
	req.Header.Add("X-API-Key", auth.apiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}

func (auth *AuthAPI) get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	auth.defaultHeaders(req)
	return client.Do(req)
}

func (auth *AuthAPI) path(p string) string {
	return fmt.Sprintf("%s%s", auth.url, p)
}

func (auth *AuthAPI) post(url string, data map[string]interface{}) (resp *http.Response, err error) {
	body, err := toJSON(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	auth.defaultHeaders(req)
	return client.Do(req)
}

func toJSON(data map[string]interface{}) (io.Reader, error) {
	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return io.Reader(bytes.NewReader(buf)), nil
}
