package main

import (
	"context"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func handlerTest(t *testing.T) (context.Context, *Handler) {
	clearDatabase()
	t.Cleanup(clearDatabase)
	return context.Background(), NewHandler(nil, nil)
}

func clearDatabase() {
	flights = []*api.Flight{}
	airports = []*api.Airport{}
}

func ptrTo[T any](v T) *T {
	return &v
}
