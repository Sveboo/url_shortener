// Package app provides entities for URL shortening
package app

import (
	"context"
	"fmt"
	"math/rand"
	"shortener/internal/storage"

	"github.com/jxskiss/base62"
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
	s       storage.Storager // underlying storage
	baseUrl string           // current host URI
}

func NewUrlShortener(storage storage.Storager, baseUrl string) *UrlShortener {
	return &UrlShortener{
		s:       storage,
		baseUrl: baseUrl,
	}
}

// CreateUrl creates a unique hash for url and writes it to storage
func (us UrlShortener) CreateUrl(ctx context.Context, url string) (string, error) {
	urlHash := Hash(url)

	for {
		if _, err := us.s.Read(ctx, urlHash); err == nil {
			urlHash = Hash(url)
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

func Hash(url string) string {
	randInt := rand.Uint64()
	shortUrl := base62.EncodeToString(base62.FormatInt(int64(randInt)))
	return shortUrl
}
