package httpserver

import (
	"context"
	"fmt"
	"shortener/internal/errs"
	"shortener/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestShortenUrl(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	s := mocks.NewMockShortener(ctr)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client := getTestClient(ctx, s)
	someErr := fmt.Errorf("some error")

	tests := []struct {
		url          string
		hashUrl      string
		errCreateUrl error
		errJson      string
	}{
		{
			url:          "http://example1.com",
			hashUrl:      "http://localhost:18080/example",
			errCreateUrl: nil,
			errJson:      "",
		},
		{
			url:          "example.com",
			hashUrl:      "",
			errCreateUrl: nil,
			errJson:      errs.ErrInvalidUrl.Error(),
		},
		{
			url:          "http://example2.com",
			hashUrl:      "",
			errCreateUrl: someErr,
			errJson:      someErr.Error(),
		},
	}

	for _, test := range tests {
		s.EXPECT().CreateUrl(ctx, test.url).Return(test.hashUrl, test.errCreateUrl).AnyTimes()
		resp, err := client.shortenUrl(test.url)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, test.errJson, resp.Err)
		assert.Equal(t, test.hashUrl, resp.Url)
	}

}

func TestGetFullUrl(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	s := mocks.NewMockShortener(ctr)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client := getTestClient(ctx, s)

	tests := []struct {
		shortUrl  string
		url       string
		errGetUrl error
		errJson   string
	}{
		{
			shortUrl:  "example",
			url:       "http://example",
			errGetUrl: nil,
			errJson:   "",
		},
		{
			shortUrl:  "ex",
			url:       "",
			errGetUrl: errs.ErrUrlNotFound,
			errJson:   errs.ErrUrlNotFound.Error(),
		},
	}

	for _, test := range tests {
		s.EXPECT().GetUrl(ctx, test.shortUrl).Return(test.url, test.errGetUrl)
		resp, err := client.getFullUrl(test.shortUrl)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, test.errJson, resp.Err)
		assert.Equal(t, test.url, resp.Url)
	}
}
