package main

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stellora/airline/api-server/api"
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

func insertAircraftT(t *testing.T, handler *Handler, airline string, registrations ...string) (ids []int) {
	t.Helper()
	ids, err := insertAircraft(context.Background(), handler, airline, "B77W", registrations...)
	if err != nil {
		t.Fatal(err)
	}
	return ids
}

func insertAirportsWithIATACodesT(t *testing.T, handler *Handler, iataCodes ...string) (ids []int) {
	t.Helper()
	ids, err := insertAirportsWithIATACodes(context.Background(), handler, iataCodes...)
	if err != nil {
		t.Fatal(err)
	}
	return ids
}

func insertAirlinesWithIATACodesT(t *testing.T, handler *Handler, iataCodes ...string) (ids []int) {
	t.Helper()
	ids, err := insertAirlinesWithIATACodes(context.Background(), handler, iataCodes...)
	if err != nil {
		t.Fatal(err)
	}
	return ids
}

func insertFlightSchedulesT(t *testing.T, handler *Handler, flightTitles ...string) (ids []int) {
	t.Helper()
	ids, err := insertFlightSchedules(context.Background(), handler, flightTitles...)
	if err != nil {
		t.Fatal(err)
	}
	return ids
}

func insertFlightScheduleT(t *testing.T, handler *Handler, startDate, endDate time.Time, daysOfWeek []int) api.FlightSchedule {
	t.Helper()
	flightSchedule, err := insertFlightSchedule(context.Background(), handler, startDate, endDate, daysOfWeek)
	if err != nil {
		t.Fatal(err)
	}
	return flightSchedule
}

func insertFlightInstanceT(t *testing.T, handler *Handler, newInstance api.CreateFlightInstanceJSONRequestBody) api.FlightInstance {
	t.Helper()
	instance, err := insertFlightInstance(context.Background(), handler, newInstance)
	if err != nil {
		t.Fatal(err)
	}
	return instance
}

func setNotesForFlightInstance(t *testing.T, handler *Handler, id int, notes string) {
	t.Helper()
	if _, err := handler.UpdateFlightInstance(context.Background(), api.UpdateFlightInstanceRequestObject{
		Id: id,
		Body: &api.UpdateFlightInstanceJSONRequestBody{
			Notes: ptrTo(notes),
		},
	}); err != nil {
		t.Fatal(err)
	}
}

func assertEqual(t *testing.T, got any, want any) {
	t.Helper()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
