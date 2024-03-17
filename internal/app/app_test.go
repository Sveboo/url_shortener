package app

import (
	"context"
	"fmt"
	"shortener/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func hashFunc(i int64) string {
	return fmt.Sprint(i)
}
func TestGetUrl(t *testing.T) {

	ctr := gomock.NewController(t)
	defer ctr.Finish()
	s := mocks.NewMockStorager(ctr)
	hash := "some_hash"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		wantUrl string
		wantErr error
	}{
		{
			wantUrl: "",
			wantErr: fmt.Errorf("some_error"),
		},
		{
			wantUrl: "http://example",
			wantErr: nil,
		},
	}
	us := NewUrlShortener(s, "http://localhost:8080", func() int64 { return int64(1) }, hashFunc)
	for _, test := range tests {
		s.EXPECT().Read(ctx, hash).Return(test.wantUrl, test.wantErr)
		got, err := us.GetUrl(ctx, hash)
		assert.Equal(t, test.wantUrl, got)
		assert.Equal(t, test.wantErr, err)
	}
}

func TestCreateUrl(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	s := mocks.NewMockStorager(ctr)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	timesCalled := 0
	randInt := func() int64 {
		timesCalled++
		return int64(timesCalled)
	}

	us := NewUrlShortener(s, "http://localhost:8080", randInt, hashFunc)
	url := "http://example"

	// read "" err write nil
	// read "" nil hash write nil
	// read "" err write err

	rErr := fmt.Errorf("some_error")
	tests := []struct {
		wantShortUrl string
		wantReadErr  error
		wantWriteErr error
	}{
		{
			wantShortUrl: "",
			wantReadErr:  rErr,
			wantWriteErr: nil,
		},
		{
			wantShortUrl: "",
			wantReadErr:  rErr,
			wantWriteErr: rErr,
		},
		{
			wantShortUrl: "",
			wantReadErr:  nil,
			wantWriteErr: nil,
		},
	}

	for _, test := range tests {
		timesCalled = 0
		hash := hashFunc(1)
		hashRetry := hash
		if test.wantReadErr == nil {
			hashRetry = hashFunc(2)
			s.EXPECT().Read(ctx, hashRetry).Return(test.wantShortUrl, rErr)
		}
		s.EXPECT().Read(ctx, hash).Return(url, test.wantReadErr)

		s.EXPECT().Write(ctx, hashRetry, url).Return(test.wantWriteErr)
		got, err := us.CreateUrl(ctx, url)
		var expected string
		if test.wantWriteErr != nil {
			expected = ""
		} else {
			expected = fmt.Sprintf("%s/%s", us.baseUrl, hashRetry)
		}
		assert.Equal(t, expected, got)
		assert.Equal(t, test.wantWriteErr, err)
	}

}
