package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
	"github.com/stellora/airline/api-server/extdata"
)

var (
	addr       = flag.String("addr", "localhost:"+defaultListenPort(), "HTTP listen address")
	dbFile     = flag.String("db", "airline.db", "database file path (sqlite3)")
	sampleData = flag.Bool("sampledata", false, "use sample data")
)

func defaultListenPort() string {
	if portStr := os.Getenv("PORT"); portStr != "" {
		return portStr
	}
	return "8080"
}

func main() {
	flag.Parse()

	ctx := context.Background()

	db, queries, err := db.Open(ctx, *dbFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := extdata.LoadAirports(); err != nil {
		log.Fatal(err)
	}

	handler := NewHandler(db, queries)

	if *sampleData {
		if err := insertSampleData(ctx, handler); err != nil {
			log.Fatal(err)
		}
		log.Printf("Sample data added to %s. Exiting.", *dbFile)
		os.Exit(0)
	}

	server := api.HandlerWithOptions(
		api.NewStrictHandler(handler, nil),
		api.StdHTTPServerOptions{},
	)

	log.Printf("Starting server on %s", *addr)
	if err := http.ListenAndServe(*addr, server); err != nil {
		log.Fatal(err)
	}
}
