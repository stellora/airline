package main

import (
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestParseRoute(t *testing.T) {
	tests := []struct {
		route                       string
		wantOrigin, wantDestination string
		wantErr                     bool
	}{
		{
			route:           "AAA-BBB",
			wantOrigin:      "AAA",
			wantDestination: "BBB",
		},
		{
			route:   "AAA--BBB",
			wantErr: true,
		},
		{
			route:   "-BBB-CCC",
			wantErr: true,
		},
		{
			route:   "   AAA-BBB",
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.route, func(t *testing.T) {
			gotOrigin, gotDestination, err := parseRoute(test.route)
			if gotOrigin != test.wantOrigin {
				t.Errorf("%q: got origin %q, want %q", test.route, gotOrigin, test.wantOrigin)
			}
			if gotDestination != test.wantDestination {
				t.Errorf("%q: got destination %q, want %q", test.route, gotDestination, test.wantDestination)
			}
			if test.wantErr != (err != nil) {
				t.Errorf("%q: got error %q, want %v", test.route, err, test.wantErr)
			}
		})
	}
}

func TestGetRoute(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB", "CCC")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 BBB-AAA", "XX3 AAA-BBB")

	t.Run("has flights", func(t *testing.T) {
		resp, err := handler.GetRoute(ctx, api.GetRouteRequestObject{Route: "AAA-BBB"})
		if err != nil {
			t.Fatal(err)
		}
		if got, want := resp.(api.GetRoute200JSONResponse).SchedulesCount, 2; got != want {
			t.Errorf("got SchedulesCount %d, want %d", got, want)
		}
	})

	t.Run("valid airports but no flights", func(t *testing.T) {
		resp, err := handler.GetRoute(ctx, api.GetRouteRequestObject{Route: "AAA-CCC"})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(*api.GetRoute404Response); !ok {
			t.Errorf("got %T", resp)
		}
	})

	t.Run("invalid airports", func(t *testing.T) {
		resp, err := handler.GetRoute(ctx, api.GetRouteRequestObject{Route: "AAA-XXX"})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(*api.GetRoute404Response); !ok {
			t.Errorf("got %T", resp)
		}
	})
}

func TestListRoutes(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB", "CCC")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 BBB-AAA", "XX3 AAA-CCC", "XX4 AAA-CCC")

	resp, err := handler.ListRoutes(ctx, api.ListRoutesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListRoutes200JSONResponse{
		{
			OriginAirport:      aaaAirport,
			DestinationAirport: cccAirport,
			SchedulesCount:     2,
		},
		{
			OriginAirport:      aaaAirport,
			DestinationAirport: bbbAirport,
			SchedulesCount:     1,
		},
		{
			OriginAirport:      bbbAirport,
			DestinationAirport: aaaAirport,
			SchedulesCount:     1,
		},
	}
	assertEqual(t, resp, want)
}

func TestListSchedulesByRoute(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB", "CCC")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 BBB-AAA", "XX3 AAA-BBB")

	t.Run("has flights", func(t *testing.T) {
		resp, err := handler.ListSchedulesByRoute(ctx, api.ListSchedulesByRouteRequestObject{
			Route: "AAA-BBB",
		})
		if err != nil {
			t.Fatal(err)
		}

		assertEqual(t, mapSlice(scheduleTitle, resp.(api.ListSchedulesByRoute200JSONResponse)), []string{"XX1 AAA-BBB", "XX3 AAA-BBB"})
	})

	t.Run("valid airports but no flights", func(t *testing.T) {
		resp, err := handler.ListSchedulesByRoute(ctx, api.ListSchedulesByRouteRequestObject{
			Route: "AAA-CCC",
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(api.ListSchedulesByRoute200JSONResponse); !ok {
			t.Errorf("got %T", resp)
		}
	})

	t.Run("invalid airports", func(t *testing.T) {
		resp, err := handler.ListSchedulesByRoute(ctx, api.ListSchedulesByRouteRequestObject{
			Route: "AAA-XXX",
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(*api.ListSchedulesByRoute404Response); !ok {
			t.Errorf("got %T", resp)
		}
	})
}

func TestListFlightsByRoute(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB", "CCC")
	insertAirlinesWithIATACodesT(t, handler, "XX")
	insertFlightT(t, handler, fixtureManualFlight)

	t.Run("has flights", func(t *testing.T) {
		resp, err := handler.ListFlightsByRoute(ctx, api.ListFlightsByRouteRequestObject{
			Route: "AAA-BBB",
		})
		if err != nil {
			t.Fatal(err)
		}

		assertEqual(t, mapSlice(flightDescription, resp.(api.ListFlightsByRoute200JSONResponse)), []string{
			"XX222 AAA-BBB on 2025-01-01",
		})
	})

	t.Run("valid airports but no flights", func(t *testing.T) {
		resp, err := handler.ListFlightsByRoute(ctx, api.ListFlightsByRouteRequestObject{
			Route: "AAA-CCC",
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(api.ListFlightsByRoute200JSONResponse); !ok {
			t.Errorf("got %T", resp)
		}
	})

	t.Run("invalid airports", func(t *testing.T) {
		resp, err := handler.ListFlightsByRoute(ctx, api.ListFlightsByRouteRequestObject{
			Route: "AAA-XXX",
		})
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := resp.(*api.ListFlightsByRoute404Response); !ok {
			t.Errorf("got %T", resp)
		}
	})
}
