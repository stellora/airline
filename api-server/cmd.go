package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/stellora/shop/api-server/api"
)

var (
	addr = flag.String("addr", "localhost:"+defaultListenPort(), "HTTP listen address")
)

func defaultListenPort() string {
	if portStr := os.Getenv("PORT"); portStr != "" {
		return portStr
	}
	return "8080"
}

func main() {
	flag.Parse()
	handler := NewHandler()
	server := api.HandlerWithOptions(
		api.NewStrictHandler(handler, nil),
		api.StdHTTPServerOptions{},
	)

	log.Printf("Starting server on %s", *addr)
	if err := http.ListenAndServe(*addr, server); err != nil {
		log.Fatal(err)
	}
}
