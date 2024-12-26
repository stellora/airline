package main

import (
	"context"
	"testing"

	"github.com/stellora/shop/api-server/api"
)

func handlerTest(t *testing.T) (context.Context, *Handler) {
	resetDatabase()
	t.Cleanup(resetDatabase)
	return context.Background(), NewHandler()
}

func resetDatabase() {
	products = []api.Product{}
	categories = []api.Category{}
	productCategoryMemberships = []productCategoryMembership{}
}
