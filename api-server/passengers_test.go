package main

import (
	"context"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

var samplePassenger = api.Passenger{
	Id:   1,
	Name: "John Doe",
}

var anotherPassenger = api.Passenger{
	Id:   2,
	Name: "Jane Smith",
}

func TestGetPassenger(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertPassengersWithNamesT(t, handler, "John Doe", "Jane Smith")

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetPassenger(ctx, api.GetPassengerRequestObject{
			Id: 1,
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, api.GetPassenger200JSONResponse(samplePassenger))
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetPassenger(ctx, api.GetPassengerRequestObject{
			Id: 999,
		})
		if err != nil {
			t.Fatal(err)
		}
		assertEqual(t, resp, &api.GetPassenger404Response{})
	})
}

func TestListPassengers(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertPassengersWithNamesT(t, handler, "John Doe", "Jane Smith")

	resp, err := handler.ListPassengers(ctx, api.ListPassengersRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, resp, api.ListPassengers200JSONResponse{
		samplePassenger,
		anotherPassenger,
	})
}

func TestCreatePassenger(t *testing.T) {
	ctx, handler := handlerTest(t)

	resp, err := handler.CreatePassenger(ctx, api.CreatePassengerRequestObject{
		Body: &api.CreatePassengerJSONRequestBody{
			Name: "John Doe",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, resp, api.CreatePassenger201JSONResponse(samplePassenger))
	checkPassengerNames(t, handler, []string{"John Doe"})
}

func TestUpdatePassenger(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertPassengersWithNamesT(t, handler, "John Doe")

	newName := "John Smith"
	resp, err := handler.UpdatePassenger(ctx, api.UpdatePassengerRequestObject{
		Id: 1,
		Body: &api.UpdatePassengerJSONRequestBody{
			Name: &newName,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := samplePassenger
	expected.Name = newName
	assertEqual(t, resp, api.UpdatePassenger200JSONResponse(expected))
	checkPassengerNames(t, handler, []string{"John Smith"})
}

func TestDeletePassenger(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertPassengersWithNamesT(t, handler, "John Doe", "Jane Smith")

	resp, err := handler.DeletePassenger(ctx, api.DeletePassengerRequestObject{
		Id: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, resp, api.DeletePassenger204Response{})
	checkPassengerNames(t, handler, []string{"Jane Smith"})
}

func checkPassengerNames(t *testing.T, handler *Handler, want []string) {
	resp, err := handler.ListPassengers(context.Background(), api.ListPassengersRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	passengers := resp.(api.ListPassengers200JSONResponse)
	if len(passengers) != len(want) {
		t.Fatalf("got %d passengers, want %d", len(passengers), len(want))
	}
	for i, passenger := range passengers {
		if passenger.Name != want[i] {
			t.Errorf("got %q, want %q", passenger.Name, want[i])
		}
	}
}

func insertPassengersWithNamesT(t *testing.T, handler *Handler, names ...string) {
	t.Helper()
	ctx := context.Background()
	for _, name := range names {
		_, err := handler.CreatePassenger(ctx, api.CreatePassengerRequestObject{
			Body: &api.CreatePassengerJSONRequestBody{
				Name: name,
			},
		})
		if err != nil {
			t.Fatal(err)
		}
	}
}
