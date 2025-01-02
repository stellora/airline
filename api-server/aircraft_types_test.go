package main

import (
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestListAircraftTypes(t *testing.T) {
	ctx, handler := handlerTest(t)

	if _, err := handler.ListAircraftTypes(ctx, api.ListAircraftTypesRequestObject{}); err != nil {
		t.Fatal(err)
	}
}
