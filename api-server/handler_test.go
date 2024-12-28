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

func ptrTo[T any](v T) *T {
	return &v
}

func insertAirportsWithIATACodesT(t *testing.T, queries *db.Queries, iataCodes ...string) (ids []int64) {
	t.Helper()
	ids, err := insertAirportsWithIATACodes(context.Background(), queries, iataCodes...)
	if err != nil {
		t.Fatal(err)
	}
	return ids
}

func insertFlightsT(t *testing.T, queries *db.Queries, flightTitles ...string) (ids []int64) {
	t.Helper()
	ids, err := insertFlights(context.Background(), queries, flightTitles...)
	if err != nil {
		t.Fatal(err)
	}
	return ids
}
