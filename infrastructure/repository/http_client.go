package repository

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

type HttpClientRepository struct{}

func NewHttpClient() *HttpClientRepository {
	return &HttpClientRepository{}
}

func (h *HttpClientRepository) SendRequest(method string, baseURL string, endpointPath string, data string, query map[string]string) (body []byte, statusCode int, err error) {
	client := &http.Client{}

	endpoint, err := url.Parse(baseURL)
	if err != nil {
		return nil, 0, err
	}

	path, err := url.Parse(endpointPath)
	if err != nil {
		return nil, 0, err
	}

	fullURL := endpoint.ResolveReference(path)

	req, err := http.NewRequest(method, fullURL.String(), bytes.NewBufferString(data))
	if err != nil {
		return nil, 0, err
	}

	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return responseBody, resp.StatusCode, nil
}
