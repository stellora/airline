package main

import (
	"context"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestGetAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")

	t.Run("exists", func(t *testing.T) {
		t.Run("by ID", func(t *testing.T) {
			resp, err := handler.GetAirport(ctx, api.GetAirportRequestObject{
				AirportSpec: api.NewAirportSpec(1, ""),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, api.GetAirport200JSONResponse{
				Id:       1,
				IataCode: "AAA",
			})
		})
		t.Run("by IATA code", func(t *testing.T) {
			resp, err := handler.GetAirport(ctx, api.GetAirportRequestObject{
				AirportSpec: api.NewAirportSpec(0, "AAA"),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, api.GetAirport200JSONResponse{
				Id:       1,
				IataCode: "AAA",
			})
		})
	})
	t.Run("does not exist", func(t *testing.T) {
		t.Run("by ID", func(t *testing.T) {
			resp, err := handler.GetAirport(ctx, api.GetAirportRequestObject{
				AirportSpec: api.NewAirportSpec(999, ""),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, &api.GetAirport404Response{})
		})
		t.Run("by IATA code", func(t *testing.T) {
			resp, err := handler.GetAirport(ctx, api.GetAirportRequestObject{
				AirportSpec: api.NewAirportSpec(0, "ZZ"),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, &api.GetAirport404Response{})
		})
	})
}

func TestListAirports(t *testing.T) {
	ctx, handler := handlerTest(t)
	ids := insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")

	resp, err := handler.ListAirports(ctx, api.ListAirportsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListAirports200JSONResponse{
		api.Airport{Id: int(ids[0]), IataCode: "AAA"},
		api.Airport{Id: int(ids[1]), IataCode: "BBB"},
	}
	assertEqual(t, resp, want)
}

func TestCreateAirport(t *testing.T) {
	ctx, handler := handlerTest(t)

	resp, err := handler.CreateAirport(ctx, api.CreateAirportRequestObject{
		Body: &api.CreateAirportJSONRequestBody{
			IataCode: "AAA",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.CreateAirport201JSONResponse{
		Id:       1,
		IataCode: "AAA",
	}
	assertEqual(t, resp, want)

	checkAirportIATACodes(t, handler, []string{"AAA"})
}

func TestDeleteAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")

	resp, err := handler.DeleteAirport(ctx, api.DeleteAirportRequestObject{
		AirportSpec: api.NewAirportSpec(1, ""),
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAirport204Response{}
	assertEqual(t, resp, want)

	checkAirportIATACodes(t, handler, []string{"BBB"})
}

func TestDeleteAllAirports(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB")

	resp, err := handler.DeleteAllAirports(ctx, api.DeleteAllAirportsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAllAirports204Response{}
	assertEqual(t, resp, want)

	checkAirportIATACodes(t, handler, []string{})
}

func checkAirportIATACodes(t *testing.T, handler *Handler, want []string) {
	resp, err := handler.ListAirports(context.Background(), api.ListAirportsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	airports := resp.(api.ListAirports200JSONResponse)
	if len(airports) != len(want) {
		t.Fatalf("got %d airports, want %d", len(airports), len(want))
	}
	for i, airport := range airports {
		if airport.IataCode != want[i] {
			t.Errorf("got %q, want %q", airport.IataCode, want[i])
		}
	}
}
