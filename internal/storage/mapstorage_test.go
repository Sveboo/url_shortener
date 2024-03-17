package storage

import (
	"context"
	"shortener/internal/errs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapstorageRead(t *testing.T) {
	hashUrl := "some_hash"
	fullUrl := "full_url"
	tests := []struct {
		hashUrl     string
		wantFullUrl string
		wantErr     error
	}{
		{
			hashUrl:     hashUrl,
			wantFullUrl: fullUrl,
			wantErr:     nil,
		},
		{
			hashUrl:     "not_existing",
			wantFullUrl: "",
			wantErr:     errs.ErrUrlNotFound,
		},
	}

	m := NewMapStorage()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	m.Write(ctx, hashUrl, fullUrl)
	for _, test := range tests {
		got, err := m.Read(ctx, test.hashUrl)
		assert.Equal(t, got, test.wantFullUrl)
		assert.Equal(t, test.wantErr, err)
	}

}

func TestMapstorageWrite(t *testing.T) {
	m := NewMapStorage()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hashUrl := "some_hash"
	fullUrl := "full_url"
	got, err := m.Read(ctx, hashUrl)
	assert.Equal(t, "", got)
	assert.Equal(t, errs.ErrUrlNotFound, err)

	m.Write(ctx, hashUrl, fullUrl)
	got, err = m.Read(ctx, hashUrl)
	assert.Equal(t, fullUrl, got)
	assert.Equal(t, nil, err)
}
