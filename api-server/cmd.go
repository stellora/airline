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

func insertSampleData(ctx context.Context, handler *Handler) error {
	iataCodes := []string{"SFO", "SIN", "NRT", "HND", "EWR", "LAX", "DEN", "ORD", "AMS", "LHR", "HKG", "SYD", "HNL", "LIH", "SJC", "STS", "SEA", "FRA", "MUC", "DXB", "TLV", "IST", "DOH", "DEL", "BOM", "KIX", "MEL", "JNB", "CPT", "EZE", "MEX", "CUN", "MSP", "IAD", "DCA", "DFW"}
	if _, err := insertAirportsWithIATACodes(ctx, handler, iataCodes...); err != nil {
		return err
	}

	flightTitles := []string{
		"UA1 SFO-SIN", "UA2 SIN-SFO",
		"UA2168 SFO-EWR", "UA1054 SFO-EWR", "UA2460 SFO-EWR",
		"UA2855 EWR-SFO", "UA598 EWR-SFO",
		"UA1684 SFO-LIH",
		"UA863 SFO-SYD", "UA830 SYD-SFO",
		"UA5703 SFO-MSP", "UA5297 MSP-SFO",
		"UA14 EWR-LHR", "UA110 EWR-LHR", "UA59 FRA-SFO",
		"KL1823 AMS-FRA", "BA430 LHR-AMS",
	}
	if _, err := insertFlights(ctx, handler, flightTitles...); err != nil {
		return err
	}

	return nil
}
