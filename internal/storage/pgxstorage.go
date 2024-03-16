// Package storage provides types to store data
package storage

import (
	"context"
	"sync"

	"shortener/internal/errs"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PgxStorage provides methods to communicate with Postgres
type PgxStorage struct {
	db *pgxpool.Pool // conntection pool
}

// make sure that the database connection will only be established once per our application lifetime
var pgOnce sync.Once

func NewPgxStorage(ctx context.Context, dbUrl string) (*PgxStorage, error) {
	var pgInstance PgxStorage
	var connErr error
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, dbUrl)
		if err != nil {
			connErr = err
			return
		}
		pgInstance = PgxStorage{db}
	})

	if connErr != nil {
		return nil, connErr
	}
	return &pgInstance, nil
}

// Ping checks connections
func (ps *PgxStorage) Ping(ctx context.Context) error {
	return ps.db.Ping(ctx)
}

// Close closes connection
func (ps *PgxStorage) Close() {
	ps.db.Close()
}

// Read returns origin url by urlHash
func (ps PgxStorage) Read(ctx context.Context, urlHash string) (string, error) {
	query := `SELECT url FROM "urls" WHERE url_hash = $1 LIMIT 1`
	row := ps.db.QueryRow(ctx, query, urlHash)
	var url string
	row.Scan(&url)
	if url == "" {
		return url, errs.ErrUrlNotFound
	}

	return url, nil
}

// Write writes urlHash and url pair
func (ps PgxStorage) Write(ctx context.Context, urlHash string, url string) error {
	query := `INSERT INTO urls (url_hash, url) VALUES ($1, $2)`
	_, err := ps.db.Exec(ctx, query, urlHash, url)

	if err != nil {
		return err
	}

	return nil
}
