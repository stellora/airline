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
