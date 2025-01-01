package main

import (
	"context"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestGetAirline(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")

	t.Run("exists", func(t *testing.T) {
		t.Run("by ID", func(t *testing.T) {
			resp, err := handler.GetAirline(ctx, api.GetAirlineRequestObject{
				AirlineSpec: newAirlineSpec(1, ""),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, api.GetAirline200JSONResponse{
				Id:       1,
				IataCode: "XX",
			})
		})
		t.Run("by IATA code", func(t *testing.T) {
			resp, err := handler.GetAirline(ctx, api.GetAirlineRequestObject{
				AirlineSpec: newAirlineSpec(0, "XX"),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, api.GetAirline200JSONResponse{
				Id:       1,
				IataCode: "XX",
			})
		})
	})
	t.Run("does not exist", func(t *testing.T) {
		t.Run("by ID", func(t *testing.T) {
			resp, err := handler.GetAirline(ctx, api.GetAirlineRequestObject{
				AirlineSpec: newAirlineSpec(999, ""),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, &api.GetAirline404Response{})
		})
		t.Run("by IATA code", func(t *testing.T) {
			resp, err := handler.GetAirline(ctx, api.GetAirlineRequestObject{
				AirlineSpec: newAirlineSpec(0, "ZZ"),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, &api.GetAirline404Response{})
		})
	})
}

func TestListAirlines(t *testing.T) {
	ctx, handler := handlerTest(t)
	ids := insertAirlinesWithIATACodesT(t, handler, "XX", "YY")

	resp, err := handler.ListAirlines(ctx, api.ListAirlinesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListAirlines200JSONResponse{
		api.Airline{Id: int(ids[0]), IataCode: "XX"},
		api.Airline{Id: int(ids[1]), IataCode: "YY"},
	}
	assertEqual(t, resp, want)
}

func TestCreateAirline(t *testing.T) {
	ctx, handler := handlerTest(t)

	resp, err := handler.CreateAirline(ctx, api.CreateAirlineRequestObject{
		Body: &api.CreateAirlineJSONRequestBody{
			IataCode: "XX",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.CreateAirline201JSONResponse{
		Id:       1,
		IataCode: "XX",
	}
	assertEqual(t, resp, want)

	checkAirlineIATACodes(t, handler, []string{"XX"})
}

func TestDeleteAirline(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")

	resp, err := handler.DeleteAirline(ctx, api.DeleteAirlineRequestObject{
		AirlineSpec: newAirlineSpec(1, ""),
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAirline204Response{}
	assertEqual(t, resp, want)

	checkAirlineIATACodes(t, handler, []string{"YY"})
}

func TestDeleteAllAirlines(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")

	resp, err := handler.DeleteAllAirlines(ctx, api.DeleteAllAirlinesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAllAirlines204Response{}
	assertEqual(t, resp, want)

	checkAirlineIATACodes(t, handler, []string{})
}

func checkAirlineIATACodes(t *testing.T, handler *Handler, want []string) {
	resp, err := handler.ListAirlines(context.Background(), api.ListAirlinesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	airlines := resp.(api.ListAirlines200JSONResponse)
	if len(airlines) != len(want) {
		t.Fatalf("got %d airlines, want %d", len(airlines), len(want))
	}
	for i, airline := range airlines {
		if airline.IataCode != want[i] {
			t.Errorf("got %q, want %q", airline.IataCode, want[i])
		}
	}
}
