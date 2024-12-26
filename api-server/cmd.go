package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/stellora/shop/api-server/api"
	"github.com/stellora/shop/api-server/db"
)

var (
	addr   = flag.String("addr", "localhost:"+defaultListenPort(), "HTTP listen address")
	dbFile = flag.String("db", "shop.db", "database file path (sqlite3)")
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

	handler := NewHandler(db, queries)
	server := api.HandlerWithOptions(
		api.NewStrictHandler(handler, nil),
		api.StdHTTPServerOptions{},
	)

	log.Printf("Starting server on %s", *addr)
	if err := http.ListenAndServe(*addr, server); err != nil {
		log.Fatal(err)
	}
}
