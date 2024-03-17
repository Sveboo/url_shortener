package httpserver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"shortener/internal/app"
	"shortener/internal/storage"

	"github.com/jxskiss/base62"
)

type UserData struct {
	Url string `json:"url"`
}

type UserUrlResponse struct {
	Url string `json:"url"`
	Err string `json:"error"`
}

var (
	errBadRequest          = fmt.Errorf("bad request")
	errIntenal             = fmt.Errorf("internal service error")
	errUnprocessableEntity = fmt.Errorf("unprocessable entity")
)

type testClient struct {
	client  *http.Client
	baseURL string
}

func hash(randInt int64) string {
	shortUrl := base62.EncodeToString(base62.FormatInt(randInt))
	return shortUrl
}

func getTestClient() *testClient {
	ctx := context.Background()
	server := newHTTPServer(ctx, ":18080", app.NewUrlShortener(storage.NewMapStorage(), "http://localhost:8080", rand.Int63, hash))
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

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest {
			return errBadRequest
		}
		if resp.StatusCode == http.StatusInternalServerError {
			return errIntenal
		}
		if resp.StatusCode == http.StatusUnprocessableEntity {
			return errUnprocessableEntity
		}
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

func (tc *testClient) createUrl(url string) (UserUrlResponse, error) {
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
