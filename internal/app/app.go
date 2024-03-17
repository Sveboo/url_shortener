// Package app provides entities for URL shortening
package app

import (
	"context"
	"fmt"
	"shortener/internal/storage"
)

// Shortener provides functionality for URL shortening and storage interaction
type Shortener interface {
	// CreateUrl is responsible for shortening a URL and writing it to a storage.
	// Returns a short URL and error if it occured
	CreateUrl(context.Context, string) (string, error)
	// GetUrl returns an origin URL by short form and error if it occured
	GetUrl(context.Context, string) (string, error)
}

// UrlShortener implements Shortener interface
type UrlShortener struct {
	s        storage.Storager // underlying storage
	baseUrl  string           // current host URI
	randInt  func() int64
	hashFunc func(int64) string
}

func NewUrlShortener(storage storage.Storager, baseUrl string, randInt func() int64, hashFunc func(int64) string) *UrlShortener {
	return &UrlShortener{
		s:        storage,
		baseUrl:  baseUrl,
		randInt:  randInt,
		hashFunc: hashFunc,
	}
}

// CreateUrl creates a unique hash for url and writes it to storage
func (us UrlShortener) CreateUrl(ctx context.Context, url string) (string, error) {
	urlHash := us.hashFunc(us.randInt())

	for {
		if _, err := us.s.Read(ctx, urlHash); err == nil {
			urlHash = us.hashFunc(us.randInt())
		} else {
			break
		}
	}

	err := us.s.Write(ctx, urlHash, url)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", us.baseUrl, urlHash), nil
}

// GetUrl converts shortUrl to an origin URL and returns it
func (us UrlShortener) GetUrl(ctx context.Context, shortUrl string) (string, error) {
	url, err := us.s.Read(ctx, shortUrl)
	if err != nil {
		return "", err
	}
	return url, nil
}
