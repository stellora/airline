package main

import (
	"context"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func handlerTest(t *testing.T) (context.Context, *Handler) {
	resetDatabase()
	t.Cleanup(resetDatabase)
	return context.Background(), NewHandler()
}

func resetDatabase() {
	flights = []api.Flight{}
	airports = []api.Airport{}
	flightAirportMemberships = []flightAirportMembership{}
}
