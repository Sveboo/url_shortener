package storage

import "context"

// Storager provides functionality to store data
type Storager interface {
	// Write writes data to the storage
	Write(context.Context, string, string) error
	// Read reads data from the storage
	Read(context.Context, string) (string, error)
}
