package pkgclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type HttpClient interface {
	Post(url string, body any) (*HttpResponse, error)
	Get(url string) (*HttpResponse, error)
	Put(url string, body any) (*HttpResponse, error)
	Delete(url string) (*HttpResponse, error)
	Patch(url string, body any) (*HttpResponse, error)
}

type httpClient struct {
	client *http.Client
}

func NewHttpClient() *httpClient {
	return &httpClient{client: &http.Client{}}
}

func (c *httpClient) Post(url string, body any) (*HttpResponse, error) {
	return c.formaRequest("POST", url, body)
}

func (c *httpClient) Get(url string) (*HttpResponse, error) {
	return c.formaRequest("GET", url, nil)
}

func (c *httpClient) Put(url string, body any) (*HttpResponse, error) {
	return c.formaRequest("PUT", url, body)
}

func (c *httpClient) Delete(url string) (*HttpResponse, error) {
	return c.formaRequest("DELETE", url, nil)
}

func (c *httpClient) Patch(url string, body any) (*HttpResponse, error) {
	return c.formaRequest("PATCH", url, body)
}

func (c *httpClient) formaRequest(method string, url string, body any) (*HttpResponse, error) {
	ctx := context.Background()
	parsedBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(parsedBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &HttpResponse{StatusCode: res.StatusCode, Body: bodyBytes}, nil
}
