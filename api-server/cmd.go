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

	if *sampleData {
		if err := insertSampleData(ctx, queries); err != nil {
			log.Fatal(err)
		}
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

func insertSampleData(ctx context.Context, queries *db.Queries) error {
	iataCodes := []string{"SFO", "SIN", "NRT", "HND", "EWR", "LAX", "DEN", "ORD", "AMS", "LHR", "HKG", "SYD", "HNL", "LIH", "SJC", "STS", "SEA", "FRA", "MUC", "DXB", "TLV", "IST", "DOH", "DEL", "BOM", "KIX", "MEL", "JNB", "CPT", "EZE", "MEX", "CUN", "MSP", "IAD", "DCA"}
	if _, err := insertAirportsWithIATACodes(ctx, queries, iataCodes...); err != nil {
		return err
	}

	flightTitles := []string{
		"UA1 SFO-SIN", "UA2 SIN-SFO",
		"UA2168 SFO-EWR", "UA1054 SFO-EWR",
		"UA2855 EWR-SFO", "UA598 EWR-SFO",
		"UA1684 SFO-LIH",
	}
	if _, err := insertFlights(ctx, queries, flightTitles...); err != nil {
		return err
	}

	return nil
}
