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
	insertFlightSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 BBB-AAA", "XX3 AAA-BBB")

	t.Run("has flights", func(t *testing.T) {
		resp, err := handler.GetRoute(ctx, api.GetRouteRequestObject{Route: "AAA-BBB"})
		if err != nil {
			t.Fatal(err)
		}
		if got, want := resp.(api.GetRoute200JSONResponse).FlightSchedulesCount, 2; got != want {
			t.Errorf("got FlightSchedulesCount %d, want %d", got, want)
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
	insertFlightSchedulesT(t, handler, "XX1 AAA-BBB", "XX2 BBB-AAA", "XX3 AAA-CCC", "XX4 AAA-CCC")

	resp, err := handler.ListRoutes(ctx, api.ListRoutesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListRoutes200JSONResponse{
		{
			OriginAirport:        aaaAirport,
			DestinationAirport:   cccAirport,
			FlightSchedulesCount: 2,
		},
		{
			OriginAirport:        aaaAirport,
			DestinationAirport:   bbbAirport,
			FlightSchedulesCount: 1,
		},
		{
			OriginAirport:        bbbAirport,
			DestinationAirport:   aaaAirport,
			FlightSchedulesCount: 1,
		},
	}
	assertEqual(t, resp, want)
}
