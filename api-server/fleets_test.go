package main

import (
	"context"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestGetFleet(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX")
	fleet := insertFleetT(t, handler, "XX", "F100", "Test Fleet")

	want := api.GetFleet200JSONResponse{
		Id: 1,
		Airline: api.Airline{
			Id:       1,
			IataCode: "XX",
		},
		Code:        "F100",
		Description: "Test Fleet",
	}

	t.Run("exists", func(t *testing.T) {
		t.Run("by ID", func(t *testing.T) {
			resp, err := handler.GetFleet(ctx, api.GetFleetRequestObject{
				AirlineSpec: api.NewAirlineSpec(0, "XX"),
				FleetSpec:   api.NewFleetSpec(int(fleet), ""),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, want)
		})

		t.Run("by code", func(t *testing.T) {
			resp, err := handler.GetFleet(ctx, api.GetFleetRequestObject{
				AirlineSpec: api.NewAirlineSpec(0, "XX"),
				FleetSpec:   api.NewFleetSpec(0, "F100"),
			})
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, resp, want)
		})
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetFleet(ctx, api.GetFleetRequestObject{
			AirlineSpec: api.NewAirlineSpec(0, "XX"),
			FleetSpec:   api.NewFleetSpec(999, ""),
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, &api.GetFleet404Response{})
	})
}

func TestListFleetsByAirline(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX")
	fleet1 := insertFleetT(t, handler, "XX", "F100", "Fleet One")
	fleet2 := insertFleetT(t, handler, "XX", "F200", "Fleet Two")

	resp, err := handler.ListFleetsByAirline(ctx, api.ListFleetsByAirlineRequestObject{
		AirlineSpec: api.NewAirlineSpec(0, "XX"),
	})
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(api.ListFleetsByAirline200JSONResponse)
	assertEqual(t, len(got), 2)
	assertEqual(t, got[0].Id, int(fleet1))
	assertEqual(t, got[1].Id, int(fleet2))
}

func TestCreateFleet(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX")

	resp, err := handler.CreateFleet(ctx, api.CreateFleetRequestObject{
		AirlineSpec: api.NewAirlineSpec(0, "XX"),
		Body: &api.CreateFleetJSONRequestBody{
			Code:        "F100",
			Description: "New Test Fleet",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(api.CreateFleet201JSONResponse)
	assertEqual(t, got.Code, "F100")
	assertEqual(t, got.Description, "New Test Fleet")
	assertEqual(t, got.Airline.Id, 1)
}

func TestUpdateFleet(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertFleetT(t, handler, "XX", "F100", "Original Fleet")

	newCode := "F200"
	newDesc := "Updated Fleet"
	resp, err := handler.UpdateFleet(ctx, api.UpdateFleetRequestObject{
		AirlineSpec: api.NewAirlineSpec(0, "XX"),
		FleetSpec:   api.NewFleetSpec(0, "F100"),
		Body: &api.UpdateFleetJSONRequestBody{
			Code:        &newCode,
			Description: &newDesc,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	got := resp.(api.UpdateFleet200JSONResponse)
	assertEqual(t, got.Code, newCode)
	assertEqual(t, got.Description, newDesc)
}

func TestDeleteFleet(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertFleetT(t, handler, "XX", "F100", "Test Fleet")

	resp, err := handler.DeleteFleet(ctx, api.DeleteFleetRequestObject{
		AirlineSpec: api.NewAirlineSpec(0, "XX"),
		FleetSpec:   api.NewFleetSpec(0, "F100"),
	})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, resp, api.DeleteFleet204Response{})
}

func insertFleetT(t *testing.T, handler *Handler, airlineIATACode, code string, description string) int64 {
	resp, err := handler.CreateFleet(context.Background(), api.CreateFleetRequestObject{
		AirlineSpec: api.NewAirlineSpec(0, airlineIATACode),
		Body: &api.CreateFleetJSONRequestBody{
			Code:        code,
			Description: description,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	created := resp.(api.CreateFleet201JSONResponse)
	return int64(created.Id)
}
