package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestGetAirline(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetAirline(ctx, api.GetAirlineRequestObject{
			Id: 1,
		})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetAirline200JSONResponse{
			Id:       1,
			IataCode: "XX",
		}
		if !reflect.DeepEqual(want, resp) {
			t.Errorf("got %v, want %v", resp, want)
		}
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetAirline(ctx, api.GetAirlineRequestObject{
			Id: 999,
		})
		if err != nil {
			t.Fatal(err)
		}
		if want := (&api.GetAirline404Response{}); !reflect.DeepEqual(resp, want) {
			t.Errorf("got %v, want %v", resp, want)
		}
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
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}
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
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkAirlineIATACodes(t, handler, []string{"XX"})
}

func TestDeleteAirline(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX", "YY")

	resp, err := handler.DeleteAirline(ctx, api.DeleteAirlineRequestObject{
		Id: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAirline204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

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
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

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
