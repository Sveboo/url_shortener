package storage

import (
	"context"
	"log"
	"shortener/internal/errs"
)

// MapStorage provides methods to store data in map
type MapStorage struct {
	urls map[string]string
}

func NewMapStorage() *MapStorage {
	return &MapStorage{urls: map[string]string{}}
}

// Read returns origin url by urlHash
func (ps *MapStorage) Read(ctx context.Context, hashUrl string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		if url, ok := ps.urls[hashUrl]; ok {
			return url, nil
		}
		return "", errs.ErrUrlNotFound
	}
}

// Write writes hashUrl as a key and url as a value into the map
func (ps *MapStorage) Write(ctx context.Context, hashUrl string, url string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if val, ok := ps.urls[hashUrl]; ok {
			log.Printf("warning: ambiguous mapping key %s to values %s and %s", hashUrl, val, url)
		}

		ps.urls[hashUrl] = url
		return nil
	}
}
