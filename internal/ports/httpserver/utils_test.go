package httpserver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"shortener/internal/app"
)

type UserData struct {
	Url string `json:"url"`
}

type UserUrlResponse struct {
	Url string `json:"url"`
	Err string `json:"error"`
}

type testClient struct {
	client  *http.Client
	baseURL string
}

func getTestClient(ctx context.Context, s app.Shortener) *testClient {
	server := newHTTPServer(ctx, ":18080", s)
	testServer := httptest.NewServer(server.Handler)

	return &testClient{
		client:  testServer.Client(),
		baseURL: testServer.URL,
	}
}

func (tc *testClient) getResponse(req *http.Request, out any) error {
	resp, err := tc.client.Do(req)
	if err != nil {
		return fmt.Errorf("unexpected error %w", err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response: %w", err)
	}

	err = json.Unmarshal(respBody, out)
	if err != nil {
		return fmt.Errorf("unable to unmarshal: %w", err)
	}

	return nil
}

func (tc *testClient) shortenUrl(url string) (UserUrlResponse, error) {
	body := map[string]string{
		"url": url,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return UserUrlResponse{}, fmt.Errorf("unable to marshal: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, tc.baseURL+"/", bytes.NewReader(data))

	if err != nil {
		return UserUrlResponse{}, fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")

	var response UserUrlResponse
	err = tc.getResponse(req, &response)
	if err != nil {
		return UserUrlResponse{}, err
	}

	return response, nil
}

func (tc *testClient) getFullUrl(shortUrl string) (UserUrlResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", tc.baseURL, shortUrl), nil)
	if err != nil {
		return UserUrlResponse{}, fmt.Errorf("unable to create request: %w", err)
	}

	var response UserUrlResponse
	err = tc.getResponse(req, &response)
	if err != nil {
		return UserUrlResponse{}, err
	}

	return response, nil
}
