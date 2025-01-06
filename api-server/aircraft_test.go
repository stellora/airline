package main

import (
	"context"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestGetAircraft(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertAircraftT(t, handler, "XX", "N1")

	t.Run("exists", func(t *testing.T) {
		t.Run("by ID", func(t *testing.T) {
			resp, err := handler.GetAircraft(ctx, api.GetAircraftRequestObject{
				AircraftSpec: api.NewAircraftSpec(1, ""),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, api.GetAircraft200JSONResponse{
				Id:           1,
				Registration: "N1",
				AircraftType: "B77W",
				Airline:      xxAirline,
			})
		})
		t.Run("by registration", func(t *testing.T) {
			resp, err := handler.GetAircraft(ctx, api.GetAircraftRequestObject{
				AircraftSpec: api.NewAircraftSpec(0, "N1"),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, api.GetAircraft200JSONResponse{
				Id:           1,
				Registration: "N1",
				AircraftType: "B77W",
				Airline:      xxAirline,
			})
		})
	})
	t.Run("does not exist", func(t *testing.T) {
		t.Run("by ID", func(t *testing.T) {
			resp, err := handler.GetAircraft(ctx, api.GetAircraftRequestObject{
				AircraftSpec: api.NewAircraftSpec(999, ""),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, &api.GetAircraft404Response{})
		})
		t.Run("by registration", func(t *testing.T) {
			resp, err := handler.GetAircraft(ctx, api.GetAircraftRequestObject{
				AircraftSpec: api.NewAircraftSpec(0, "ZZ"),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, &api.GetAircraft404Response{})
		})
	})
}

func TestListAircraft(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX")
	ids := insertAircraftT(t, handler, "XX", "N1", "N2")

	resp, err := handler.ListAircraft(ctx, api.ListAircraftRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListAircraft200JSONResponse{
		api.Aircraft{Id: int(ids[0]), Registration: "N1", AircraftType: "B77W", Airline: xxAirline},
		api.Aircraft{Id: int(ids[1]), Registration: "N2", AircraftType: "B77W", Airline: xxAirline},
	}
	assertEqual(t, resp, want)
}

func TestCreateAircraft(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX")

	resp, err := handler.CreateAircraft(ctx, api.CreateAircraftRequestObject{
		Body: &api.CreateAircraftJSONRequestBody{
			Registration: "N1",
			AircraftType: "B77W",
			Airline:      api.NewAirlineSpec(0, "XX"),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.CreateAircraft201JSONResponse{
		Id:           1,
		Registration: "N1",
		AircraftType: "B77W",
		Airline:      xxAirline,
	}
	assertEqual(t, resp, want)

	checkAircraftRegistrations(t, handler, []string{"N1"})
}

func TestDeleteAircraft(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertAircraftT(t, handler, "XX", "N1", "N2")

	resp, err := handler.DeleteAircraft(ctx, api.DeleteAircraftRequestObject{
		AircraftSpec: api.NewAircraftSpec(1, ""),
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAircraft204Response{}
	assertEqual(t, resp, want)

	checkAircraftRegistrations(t, handler, []string{"N2"})
}

func TestDeleteAllAircraft(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertAircraftT(t, handler, "XX", "N1")

	resp, err := handler.DeleteAllAircraft(ctx, api.DeleteAllAircraftRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteAllAircraft204Response{}
	assertEqual(t, resp, want)

	checkAircraftRegistrations(t, handler, []string{})
}

func checkAircraftRegistrations(t *testing.T, handler *Handler, want []string) {
	resp, err := handler.ListAircraft(context.Background(), api.ListAircraftRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	airlines := resp.(api.ListAircraft200JSONResponse)
	if len(airlines) != len(want) {
		t.Fatalf("got %d airlines, want %d", len(airlines), len(want))
	}
	for i, airline := range airlines {
		if airline.Registration != want[i] {
			t.Errorf("got %q, want %q", airline.Registration, want[i])
		}
	}
}
