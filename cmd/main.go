// Package main is an entry point for the entire project
//
//
// Usage:
// gofmt [flag]
//
// The flag is:
//
// -d
//
//	Store data in PostgreSQL, path is given from the environment variable PATH_TO_DB

// Without -d flag it is used an internal storage
package main

import (
	"context"
	"flag"
	"shortener/internal/app"

	"fmt"
	"log"
	"os"
	"os/signal"

	"shortener/internal/ports/httpserver"
	"shortener/internal/storage"
	"syscall"

	_ "shortener/docs"

	"golang.org/x/sync/errgroup"
)

const (
	httpPort   = ":8080"
	domainName = "http://localhost"
)

func captureSigQuit(ctx context.Context) func() error {
	return func() error {
		sigQuit := make(chan os.Signal, 1)
		signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
		signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case s := <-sigQuit:
			log.Printf("captured signal: %v\n", s)
			return fmt.Errorf("captured signal: %v ", s)
		case <-ctx.Done():
			return nil
		}
	}
}

//	@title			Url shortener documentation
//	@version		0.1
//	@description	A collection of endpoints available to communicate with url shortener

//	@contact.name	Maintainer
//	@contact.url	https://github.com/Sveboo/url_shortener
//	@contact.email	svebo3348@gmail.com
//	@license.name	MIT
//	@license.url	https://github.com/Sveboo/url_shortener/blob/main/LICENSE
//	@host			localhost:8080
//	@accept			json
//	@produce		json
//	@schemes		http
func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	// capture signals to stop working
	eg.Go(captureSigQuit(ctx))

	// get option for storaging
	var s storage.Storager
	usePostgres := flag.Bool("d", false, "use postgres sql")
	flag.Parse()

	if *usePostgres {
		var err error
		pathToDb := os.Getenv("PATH_TO_DB")
		if pathToDb == "" {
			log.Fatal("PATH_TO_DB env variable must be set")
		}
		// path := "postgresql://postgres:password@127.0.0.1:5432/urls"
		s, err = storage.NewPgxStorage(ctx, pathToDb)
		if err != nil {
			log.Fatalf("unable to create connection with SQL: %s", err.Error())
		}
	} else {
		s = storage.NewMapStorage()
	}

	us := app.NewUrlShortener(s, fmt.Sprintf("%s%s", domainName, httpPort))
	// run HTTP server
	eg.Go(httpserver.Run(ctx, us, httpPort))

	err := eg.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
}
