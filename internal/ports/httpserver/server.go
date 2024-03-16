// Package httpserver provides utilities to run http server
package httpserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"shortener/internal/app"
	"time"
)

func newHTTPServer(ctx context.Context, port string, s app.Shortener) *http.Server {
	api := api{
		ctx:       ctx,
		shortener: s,
	}
	mux := route(api)
	service := &http.Server{Addr: port, Handler: mux}
	return service
}

// Run runs Sortener app no the given port
func Run(ctx context.Context, s app.Shortener, port string) func() error {
	return func() error {
		httpServer := newHTTPServer(ctx, port, s)

		errCh := make(chan error)

		defer func() {
			shCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := httpServer.Shutdown(shCtx); err != nil {
				log.Printf("can't close http server listening on %s %s", httpServer.Addr, err.Error())
			}

			close(errCh)
		}()

		go func() {
			if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				errCh <- err
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return fmt.Errorf("http server can't listen and serve requests: %w", err)
		}
	}
}
