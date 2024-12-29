package main

import (
	"reflect"
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

func TestListRoutes(t *testing.T) {
	ctx, handler := handlerTest(t)
	insertAirportsWithIATACodesT(t, handler, "AAA", "BBB", "CCC")
	insertFlightsT(t, handler, "ST1 AAA-BBB", "ST2 BBB-AAA", "ST3 AAA-CCC", "ST4 AAA-CCC")

	resp, err := handler.ListRoutes(ctx, api.ListRoutesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListRoutes200JSONResponse{
		{
			OriginAirport:      api.Airport{Id: 1, IataCode: "AAA"},
			DestinationAirport: api.Airport{Id: 3, IataCode: "CCC"},
			FlightsCount:       2,
		},
		{
			OriginAirport:      api.Airport{Id: 1, IataCode: "AAA"},
			DestinationAirport: api.Airport{Id: 2, IataCode: "BBB"},
			FlightsCount:       1,
		},
		{
			OriginAirport:      api.Airport{Id: 2, IataCode: "BBB"},
			DestinationAirport: api.Airport{Id: 1, IataCode: "AAA"},
			FlightsCount:       1,
		},
	}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}
}
