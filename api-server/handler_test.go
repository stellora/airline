package main

import (
	"context"
	"testing"

	"github.com/stellora/airline/api-server/db"
)

func handlerTest(t *testing.T) (context.Context, *Handler) {
	ctx := context.Background()
	db, queries, err := db.Open(ctx, ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	return ctx, NewHandler(db, queries)
}

func insertAirportsWithIATACodesT(t *testing.T, handler *Handler, iataCodes ...string) (ids []int) {
	t.Helper()
	ids, err := insertAirportsWithIATACodes(context.Background(), handler, iataCodes...)
	if err != nil {
		t.Fatal(err)
	}
	return ids
}

func insertAirlinesWithIATACodesT(t *testing.T, handler *Handler, iataCodes ...string) (ids []int) {
	t.Helper()
	ids, err := insertAirlinesWithIATACodes(context.Background(), handler, iataCodes...)
	if err != nil {
		t.Fatal(err)
	}
	return ids
}

func insertFlightsT(t *testing.T, handler *Handler, flightTitles ...string) (ids []int) {
	t.Helper()
	ids, err := insertFlights(context.Background(), handler, flightTitles...)
	if err != nil {
		t.Fatal(err)
	}
	return ids
}
