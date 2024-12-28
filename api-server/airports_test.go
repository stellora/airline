package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/stellora/airline/api-server/api"
	"github.com/stellora/airline/api-server/db"
)

func insertAirportsWithIATACodes(t *testing.T, queries *db.Queries, iataCodes ...string) (ids []int64) {
	t.Helper()
	ids = make([]int64, len(iataCodes))
	for i, iataCode := range iataCodes {
		v, err := queries.CreateAirport(context.Background(), db.CreateAirportParams{IataCode: iataCode})
		if err != nil {
			t.Fatal(err)
		}
		ids[i] = v.ID
	}
	return ids
}

func TestGetAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodes(t, handler.queries, "AAA", "BBB")

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetAirport(ctx, api.GetAirportRequestObject{
			Id: 1,
		})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetAirport200JSONResponse{
			Id:       1,
			IataCode: "AAA",
		}
		if !reflect.DeepEqual(want, resp) {
			t.Errorf("got %v, want %v", resp, want)
		}
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetAirport(ctx, api.GetAirportRequestObject{
			Id: 999,
		})
		if err != nil {
			t.Fatal(err)
		}
		if want := (&api.GetAirport404Response{}); !reflect.DeepEqual(resp, want) {
			t.Errorf("got %v, want %v", resp, want)
		}
	})
}

func TestListAirports(t *testing.T) {
	ctx, handler := handlerTest(t)
	ids := insertAirportsWithIATACodes(t, handler.queries, "AAA", "BBB")

	resp, err := handler.ListAirports(ctx, api.ListAirportsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListAirports200JSONResponse{
		api.Airport{Id: int(ids[0]), IataCode: "AAA"},
		api.Airport{Id: int(ids[1]), IataCode: "BBB"},
	}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}
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

	want := api.CreateAirport201Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkAirportIATACodes(t, handler, []string{"AAA"})
}

func TestDeleteAirport(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodes(t, handler.queries, "AAA", "BBB")

	resp, err := handler.DeleteAirport(ctx, api.DeleteAirportRequestObject{
		Id: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAirport204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkAirportIATACodes(t, handler, []string{"BBB"})
}

func TestDeleteAllAirports(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodes(t, handler.queries, "AAA", "BBB")

	resp, err := handler.DeleteAllAirports(ctx, api.DeleteAllAirportsRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAllAirports204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

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
