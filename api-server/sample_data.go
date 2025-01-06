package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/stellora/airline/api-server/api"
)

func insertSampleData(ctx context.Context, handler *Handler) error {
	airports := []string{"SFO", "SIN", "NRT", "HND", "EWR", "LAX", "DEN", "ORD", "AMS", "LHR", "HKG", "SYD", "HNL", "LIH", "SJC", "STS", "SEA", "FRA", "MUC", "DXB", "TLV", "IST", "DOH", "KIX", "MEL", "JNB", "CPT", "EZE", "MEX", "CUN", "MSP", "IAD", "DCA", "DFW", "MRS", "PVR", "BOS", "FCO", "SCL", "CDG", "ICN", "PVG", "HYD", "DEL", "BOM", "BLR", "PEK"}
	if _, err := insertAirportsWithIATACodes(ctx, handler, airports...); err != nil {
		return err
	}

	airlines := map[string]string{
		"UA": "United Airlines",
		"SQ": "Singapore Airlines",
		"LH": "Lufthansa",
		"BA": "British Airways",
		"DL": "Delta Air Lines",
		"AF": "Air France",
		"KL": "KLM Royal Dutch Airlines",
		"LX": "Swiss International Air Lines",
		"EY": "Etihad Airways",
		"EK": "Emirates",
		"QR": "Qatar Airways",
		"QF": "Qantas",
		"NZ": "Air New Zealand",
	}
	if _, err := insertAirlines(ctx, handler, airlines); err != nil {
		return err
	}

	// See https://sites.google.com/site/unitedfleetsite/mainline-fleet-tracking for UA fleet data.
	ua := api.NewAirlineSpec(0, "UA")
	aircraft := []api.CreateAircraftJSONRequestBody{
		{Registration: "N801UA", AircraftType: "A320", Airline: ua},
		{Registration: "N802UA", AircraftType: "A320", Airline: ua},
		{Registration: "N803UA", AircraftType: "A320", Airline: ua},
		{Registration: "N804UA", AircraftType: "A320", Airline: ua},
		{Registration: "N805UA", AircraftType: "A320", Airline: ua},
		{Registration: "N806UA", AircraftType: "A320", Airline: ua},
		{Registration: "N807UA", AircraftType: "A320", Airline: ua},
		{Registration: "N808UA", AircraftType: "A320", Airline: ua},
		{Registration: "N809UA", AircraftType: "A320", Airline: ua},
		{Registration: "N810UA", AircraftType: "A320", Airline: ua},
		{Registration: "N811UA", AircraftType: "A320", Airline: ua},
		{Registration: "N812UA", AircraftType: "A320", Airline: ua},
		{Registration: "N813UA", AircraftType: "A320", Airline: ua},
		{Registration: "N814UA", AircraftType: "A320", Airline: ua},
		{Registration: "N815UA", AircraftType: "A320", Airline: ua},
		{Registration: "N816UA", AircraftType: "A320", Airline: ua},
		{Registration: "N817UA", AircraftType: "A320", Airline: ua},
		{Registration: "N818UA", AircraftType: "A320", Airline: ua},
		{Registration: "N819UA", AircraftType: "A320", Airline: ua},
		{Registration: "N820UA", AircraftType: "A320", Airline: ua},
		{Registration: "N821UA", AircraftType: "A320", Airline: ua},
		{Registration: "N822UA", AircraftType: "A320", Airline: ua},
		{Registration: "N823UA", AircraftType: "A320", Airline: ua},
		{Registration: "N824UA", AircraftType: "A320", Airline: ua},
		{Registration: "N825UA", AircraftType: "A320", Airline: ua},
		{Registration: "N826UA", AircraftType: "A320", Airline: ua},
		{Registration: "N827UA", AircraftType: "A320", Airline: ua},
		{Registration: "N828UA", AircraftType: "A320", Airline: ua},
		{Registration: "N829UA", AircraftType: "A320", Airline: ua},
		{Registration: "N830UA", AircraftType: "A320", Airline: ua},
		{Registration: "N831UA", AircraftType: "A320", Airline: ua},
		{Registration: "N832UA", AircraftType: "A320", Airline: ua},
		{Registration: "N833UA", AircraftType: "A320", Airline: ua},
		{Registration: "N834UA", AircraftType: "A320", Airline: ua},
		{Registration: "N835UA", AircraftType: "A320", Airline: ua},
		{Registration: "N836UA", AircraftType: "A320", Airline: ua},
		{Registration: "N837UA", AircraftType: "A320", Airline: ua},
		{Registration: "N838UA", AircraftType: "A320", Airline: ua},
		{Registration: "N839UA", AircraftType: "A320", Airline: ua},
		{Registration: "N840UA", AircraftType: "A320", Airline: ua},
		{Registration: "N841UA", AircraftType: "A320", Airline: ua},
		{Registration: "N842UA", AircraftType: "A320", Airline: ua},
		{Registration: "N843UA", AircraftType: "A320", Airline: ua},
		{Registration: "N844UA", AircraftType: "A320", Airline: ua},
		{Registration: "N845UA", AircraftType: "A320", Airline: ua},
		{Registration: "N846UA", AircraftType: "A320", Airline: ua},
		{Registration: "N847UA", AircraftType: "A320", Airline: ua},
		{Registration: "N848UA", AircraftType: "A320", Airline: ua},
		{Registration: "N849UA", AircraftType: "A320", Airline: ua},
		{Registration: "N850UA", AircraftType: "A320", Airline: ua},
		{Registration: "N851UA", AircraftType: "A320", Airline: ua},
		{Registration: "N852UA", AircraftType: "A320", Airline: ua},
		{Registration: "N853UA", AircraftType: "A320", Airline: ua},
		{Registration: "N854UA", AircraftType: "A320", Airline: ua},
		{Registration: "N855UA", AircraftType: "A320", Airline: ua},

		{Registration: "N27254", AircraftType: "B38M", Airline: ua},
		{Registration: "N27255", AircraftType: "B38M", Airline: ua},
		{Registration: "N27256", AircraftType: "B38M", Airline: ua},
		{Registration: "N27257", AircraftType: "B38M", Airline: ua},
		{Registration: "N27258", AircraftType: "B38M", Airline: ua},
		{Registration: "N27259", AircraftType: "B38M", Airline: ua},
		{Registration: "N27260", AircraftType: "B38M", Airline: ua},
		{Registration: "N27261", AircraftType: "B38M", Airline: ua},
		{Registration: "N27262", AircraftType: "B38M", Airline: ua},
		{Registration: "N27263", AircraftType: "B38M", Airline: ua},
		{Registration: "N27264", AircraftType: "B38M", Airline: ua},
		{Registration: "N27265", AircraftType: "B38M", Airline: ua},
		{Registration: "N27266", AircraftType: "B38M", Airline: ua},
		{Registration: "N27267", AircraftType: "B38M", Airline: ua},
		{Registration: "N27268", AircraftType: "B38M", Airline: ua},
		{Registration: "N27269", AircraftType: "B38M", Airline: ua},
		{Registration: "N27270", AircraftType: "B38M", Airline: ua},
		{Registration: "N27271", AircraftType: "B38M", Airline: ua},
		{Registration: "N27272", AircraftType: "B38M", Airline: ua},
		{Registration: "N27273", AircraftType: "B38M", Airline: ua},
		{Registration: "N27274", AircraftType: "B38M", Airline: ua},
		{Registration: "N47275", AircraftType: "B38M", Airline: ua},
		{Registration: "N27276", AircraftType: "B38M", Airline: ua},
		{Registration: "N37278", AircraftType: "B38M", Airline: ua},

		{Registration: "N2331U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2332U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2333U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2534U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2135U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2136U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2737U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2138U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2639U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2140U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2341U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2142U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2243U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2644U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2645U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2846U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2747U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2748U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2749U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2250U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2251U", AircraftType: "B77W", Airline: ua},
		{Registration: "N2352U", AircraftType: "B77W", Airline: ua},

		{Registration: "N27901", AircraftType: "B788", Airline: ua},
		{Registration: "N26902", AircraftType: "B788", Airline: ua},
		{Registration: "N27903", AircraftType: "B788", Airline: ua},
		{Registration: "N20904", AircraftType: "B788", Airline: ua},
		{Registration: "N45905", AircraftType: "B788", Airline: ua},
		{Registration: "N26906", AircraftType: "B788", Airline: ua},
		{Registration: "N29907", AircraftType: "B788", Airline: ua},
		{Registration: "N27908", AircraftType: "B788", Airline: ua},
		{Registration: "N26909", AircraftType: "B788", Airline: ua},
		{Registration: "N26910", AircraftType: "B788", Airline: ua},
		{Registration: "N28912", AircraftType: "B788", Airline: ua},
		{Registration: "N30913", AircraftType: "B788", Airline: ua},

		{Registration: "N38950", AircraftType: "B789", Airline: ua},
		{Registration: "N19951", AircraftType: "B789", Airline: ua},
		{Registration: "N26952", AircraftType: "B789", Airline: ua},
		{Registration: "N35953", AircraftType: "B789", Airline: ua},
		{Registration: "N13954", AircraftType: "B789", Airline: ua},
		{Registration: "N38955", AircraftType: "B789", Airline: ua},
		{Registration: "N45956", AircraftType: "B789", Airline: ua},
		{Registration: "N27957", AircraftType: "B789", Airline: ua},
		{Registration: "N27958", AircraftType: "B789", Airline: ua},
		{Registration: "N27959", AircraftType: "B789", Airline: ua},
		{Registration: "N26960", AircraftType: "B789", Airline: ua},
		{Registration: "N29961", AircraftType: "B789", Airline: ua},
		{Registration: "N36962", AircraftType: "B789", Airline: ua},
		{Registration: "N17963", AircraftType: "B789", Airline: ua},
		{Registration: "N27964", AircraftType: "B789", Airline: ua},
		{Registration: "N27965", AircraftType: "B789", Airline: ua},
		{Registration: "N26966", AircraftType: "B789", Airline: ua},
		{Registration: "N26967", AircraftType: "B789", Airline: ua},
		{Registration: "N29968", AircraftType: "B789", Airline: ua},
		{Registration: "N15969", AircraftType: "B789", Airline: ua},
		{Registration: "N26970", AircraftType: "B789", Airline: ua},
		{Registration: "N29971", AircraftType: "B789", Airline: ua},
		{Registration: "N24972", AircraftType: "B789", Airline: ua},
		{Registration: "N24973", AircraftType: "B789", Airline: ua},
		{Registration: "N24974", AircraftType: "B789", Airline: ua},
		{Registration: "N29975", AircraftType: "B789", Airline: ua},
		{Registration: "N24976", AircraftType: "B789", Airline: ua},
		{Registration: "N29977", AircraftType: "B789", Airline: ua},
		{Registration: "N29978", AircraftType: "B789", Airline: ua},
		{Registration: "N24979", AircraftType: "B789", Airline: ua},
		{Registration: "N24980", AircraftType: "B789", Airline: ua},
		{Registration: "N29981", AircraftType: "B789", Airline: ua},
		{Registration: "N25982", AircraftType: "B789", Airline: ua},
		{Registration: "N23983", AircraftType: "B789", Airline: ua},
		{Registration: "N29984", AircraftType: "B789", Airline: ua},
		{Registration: "N29985", AircraftType: "B789", Airline: ua},
		{Registration: "N19986", AircraftType: "B789", Airline: ua},
		{Registration: "N28987", AircraftType: "B789", Airline: ua},
		{Registration: "N24988", AircraftType: "B789", Airline: ua},
		{Registration: "N29989", AircraftType: "B789", Airline: ua},
		{Registration: "N24990", AircraftType: "B789", Airline: ua},

		{Registration: "N14001", AircraftType: "B78X", Airline: ua},
		{Registration: "N17002", AircraftType: "B78X", Airline: ua},
		{Registration: "N12003", AircraftType: "B78X", Airline: ua},
		{Registration: "N12004", AircraftType: "B78X", Airline: ua},
		{Registration: "N12005", AircraftType: "B78X", Airline: ua},
		{Registration: "N12006", AircraftType: "B78X", Airline: ua},
		{Registration: "N91007", AircraftType: "B78X", Airline: ua},
		{Registration: "N16008", AircraftType: "B78X", Airline: ua},
		{Registration: "N16009", AircraftType: "B78X", Airline: ua},
		{Registration: "N12010", AircraftType: "B78X", Airline: ua},
		{Registration: "N14011", AircraftType: "B78X", Airline: ua},
		{Registration: "N12012", AircraftType: "B78X", Airline: ua},
		{Registration: "N13013", AircraftType: "B78X", Airline: ua},
		{Registration: "N13014", AircraftType: "B78X", Airline: ua},
		{Registration: "N17015", AircraftType: "B78X", Airline: ua},
		{Registration: "N14016", AircraftType: "B78X", Airline: ua},
		{Registration: "N17017", AircraftType: "B78X", Airline: ua},
		{Registration: "N13018", AircraftType: "B78X", Airline: ua},
		{Registration: "N14019", AircraftType: "B78X", Airline: ua},
		{Registration: "N12020", AircraftType: "B78X", Airline: ua},
		{Registration: "N12021", AircraftType: "B78X", Airline: ua},

		{Registration: "9V-SJG", AircraftType: "A359", Airline: api.NewAirlineSpec(0, "SQ")},
		{Registration: "9V-SJI", AircraftType: "A359", Airline: api.NewAirlineSpec(0, "SQ")},
		{Registration: "9V-SNC", AircraftType: "B77W", Airline: api.NewAirlineSpec(0, "SQ")},

		{Registration: "D-ABVM", AircraftType: "B744", Airline: api.NewAirlineSpec(0, "LH")},
		{Registration: "D-AIXF", AircraftType: "A359", Airline: api.NewAirlineSpec(0, "LH")},
		{Registration: "D-ABYU", AircraftType: "B748", Airline: api.NewAirlineSpec(0, "LH")},
		{Registration: "D-AIXJ", AircraftType: "A359", Airline: api.NewAirlineSpec(0, "LH")},

		{Registration: "G-TTNR", AircraftType: "A20N", Airline: api.NewAirlineSpec(0, "BA")},

		{Registration: "PH-EXY", AircraftType: "E190", Airline: api.NewAirlineSpec(0, "KL")},
	}
	for _, a := range aircraft {
		if _, err := handler.CreateAircraft(ctx, api.CreateAircraftRequestObject{Body: &a}); err != nil {
			return err
		}
	}

	log.Println("Creating fleets...")
	fleetsByAirline := map[string][]api.CreateFleetJSONRequestBody{
		"UA": []api.CreateFleetJSONRequestBody{
			{Code: "B789", Description: "All 787-9s"},
			{Code: "B77W", Description: "All 77Ws"},
			{Code: "B738", Description: "All 737 MAX 8 and 737-800"},
			{Code: "B739", Description: "All 737 MAX 9 and 737-900"},
		},
		"LH": []api.CreateFleetJSONRequestBody{
			{Code: "B747", Description: "All 747s"},
			{Code: "A350", Description: "All A350s"},
		},
		"SQ": []api.CreateFleetJSONRequestBody{
			{Code: "B777", Description: "All 777s"},
			{Code: "B359", Description: "All 359s"},
		},
		"KL": []api.CreateFleetJSONRequestBody{{Code: "E190", Description: "All E190s"}},
		"BA": []api.CreateFleetJSONRequestBody{{Code: "A20N", Description: "All A320 neos"}},
	}
	for airlineCode, fleets := range fleetsByAirline {
		for _, fleet := range fleets {
			if _, err := handler.CreateFleet(ctx, api.CreateFleetRequestObject{
				AirlineSpec: api.NewAirlineSpec(0, airlineCode),
				Body:        &fleet,
			}); err != nil {
				return fmt.Errorf("creating fleet %s: %w", fleet.Code, err)
			}
		}
	}

	log.Println("Adding aircraft to fleets...")
	for _, aircraft := range aircraft {
		airline, _ := aircraft.Airline.AsAirlineIATACode()
		var fleetCode api.FleetCode
		switch airline {
		case "UA":
			if aircraft.AircraftType == "789" {
				fleetCode = "B789"
			} else if aircraft.AircraftType == "B77W" {
				fleetCode = "B77W"
			} else if aircraft.AircraftType == "B38M" || aircraft.AircraftType == "B738" {
				fleetCode = "B738"
			} else if aircraft.AircraftType == "B39M" || aircraft.AircraftType == "B739" {
				fleetCode = "B739"
			}
		case "LH":
			if aircraft.AircraftType == "B744" || aircraft.AircraftType == "B748" {
				fleetCode = "B747"
			} else if aircraft.AircraftType == "A359" {
				fleetCode = "A350"
			}
		case "SQ":
			if strings.HasPrefix(aircraft.AircraftType, "B77") {
				fleetCode = "B777"
			} else if aircraft.AircraftType == "A359" {
				fleetCode = "A359"
			}
		}
		if fleetCode == "" {
			continue
		}
		if _, err := handler.AddAircraftToFleet(ctx, api.AddAircraftToFleetRequestObject{
			AirlineSpec:  aircraft.Airline,
			AircraftSpec: api.NewAircraftSpec(0, aircraft.Registration),
			FleetSpec:    api.NewFleetSpec(0, fleetCode),
		}); err != nil {
			return fmt.Errorf("adding aircraft to fleet: %w", err)
		}
	}

	durationSec := func(hours, minutes int) int {
		return (hours*60 + minutes) * 60
	}
	flightSchedules := []*api.CreateFlightScheduleJSONRequestBody{
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "1",
			OriginAirport:      api.NewAirportSpec(0, "SFO"),
			DestinationAirport: api.NewAirportSpec(0, "SIN"),
			Fleet:              api.NewFleetSpec(0, "B789"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 2, 4, 6},
			DepartureTime:      "23:35",
			DurationSec:        durationSec(16, 30),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "2",
			OriginAirport:      api.NewAirportSpec(0, "SIN"),
			DestinationAirport: api.NewAirportSpec(0, "SFO"),
			Fleet:              api.NewFleetSpec(0, "B789"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{1, 3, 5, 0},
			DepartureTime:      "10:10",
			DurationSec:        durationSec(17, 30),
		},
		{
			Airline:            api.NewAirlineSpec(0, "SQ"),
			Number:             "33",
			OriginAirport:      api.NewAirportSpec(0, "SFO"),
			DestinationAirport: api.NewAirportSpec(0, "SIN"),
			Fleet:              api.NewFleetSpec(0, "A359"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{1, 3, 5, 0},
			DepartureTime:      "22:35",
			DurationSec:        durationSec(16, 30),
		},
		{
			Airline:            api.NewAirlineSpec(0, "SQ"),
			Number:             "34",
			OriginAirport:      api.NewAirportSpec(0, "SIN"),
			DestinationAirport: api.NewAirportSpec(0, "SFO"),
			Fleet:              api.NewFleetSpec(0, "A359"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{1, 3, 5, 0},
			DepartureTime:      "09:20",
			DurationSec:        durationSec(16, 30),
		},
		{
			Airline:            api.NewAirlineSpec(0, "SQ"),
			Number:             "31",
			OriginAirport:      api.NewAirportSpec(0, "SFO"),
			DestinationAirport: api.NewAirportSpec(0, "SIN"),
			Fleet:              api.NewFleetSpec(0, "A359"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 2, 4, 6},
			DepartureTime:      "20:35",
			DurationSec:        durationSec(16, 30),
		},
		{
			Airline:            api.NewAirlineSpec(0, "SQ"),
			Number:             "32",
			OriginAirport:      api.NewAirportSpec(0, "SIN"),
			DestinationAirport: api.NewAirportSpec(0, "SFO"),
			Fleet:              api.NewFleetSpec(0, "A359"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{1, 3, 5, 0},
			DepartureTime:      "07:20",
			DurationSec:        durationSec(16, 30),
		},
		{
			Airline:            api.NewAirlineSpec(0, "SQ"),
			Number:             "345",
			OriginAirport:      api.NewAirportSpec(0, "ZRH"),
			DestinationAirport: api.NewAirportSpec(0, "SIN"),
			Fleet:              api.NewFleetSpec(0, "B777"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{1, 3, 5, 0},
			DepartureTime:      "11:45",
			DurationSec:        durationSec(12, 30),
		},
		{
			Airline:            api.NewAirlineSpec(0, "SQ"),
			Number:             "346",
			OriginAirport:      api.NewAirportSpec(0, "SIN"),
			DestinationAirport: api.NewAirportSpec(0, "ZRH"),
			Fleet:              api.NewFleetSpec(0, "B777"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{1, 3, 5, 0},
			DepartureTime:      "01:30",
			DurationSec:        durationSec(12, 45),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "2168",
			OriginAirport:      api.NewAirportSpec(0, "SFO"),
			DestinationAirport: api.NewAirportSpec(0, "EWR"),
			Fleet:              api.NewFleetSpec(0, "B78X"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "13:30",
			DurationSec:        durationSec(5, 30),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "1054",
			OriginAirport:      api.NewAirportSpec(0, "SFO"),
			DestinationAirport: api.NewAirportSpec(0, "EWR"),
			Fleet:              api.NewFleetSpec(0, "B78X"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "06:30",
			DurationSec:        durationSec(5, 30),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "2460",
			OriginAirport:      api.NewAirportSpec(0, "SFO"),
			DestinationAirport: api.NewAirportSpec(0, "EWR"),
			Fleet:              api.NewFleetSpec(0, "B78X"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "22:30",
			DurationSec:        durationSec(5, 30),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "2855",
			OriginAirport:      api.NewAirportSpec(0, "EWR"),
			DestinationAirport: api.NewAirportSpec(0, "SFO"),
			Fleet:              api.NewFleetSpec(0, "B78X"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "17:30",
			DurationSec:        durationSec(5, 30),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "598",
			OriginAirport:      api.NewAirportSpec(0, "EWR"),
			DestinationAirport: api.NewAirportSpec(0, "SFO"),
			Fleet:              api.NewFleetSpec(0, "B78X"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "08:30",
			DurationSec:        durationSec(5, 55),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "1579",
			OriginAirport:      api.NewAirportSpec(0, "EWR"),
			DestinationAirport: api.NewAirportSpec(0, "SFO"),
			Fleet:              api.NewFleetSpec(0, "B78X"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "19:30",
			DurationSec:        durationSec(5, 55),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "1684",
			OriginAirport:      api.NewAirportSpec(0, "SFO"),
			DestinationAirport: api.NewAirportSpec(0, "LIH"),
			Fleet:              api.NewFleetSpec(0, "B738-ETOPS"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 2, 4, 6},
			DepartureTime:      "23:35",
			DurationSec:        durationSec(5, 45),
		},

		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "1563",
			OriginAirport:      api.NewAirportSpec(0, "LIH"),
			DestinationAirport: api.NewAirportSpec(0, "SFO"),
			Fleet:              api.NewFleetSpec(0, "B738-ETOPS"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 2, 4, 6},
			DepartureTime:      "07:30",
			DurationSec:        durationSec(5, 15),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "1111",
			OriginAirport:      api.NewAirportSpec(0, "SFO"),
			DestinationAirport: api.NewAirportSpec(0, "LIH"),
			Fleet:              api.NewFleetSpec(0, "B738-ETOPS"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "11:00",
			DurationSec:        durationSec(5, 45),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "863",
			OriginAirport:      api.NewAirportSpec(0, "SFO"),
			DestinationAirport: api.NewAirportSpec(0, "SYD"),
			Fleet:              api.NewFleetSpec(0, "B789"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "22:45",
			DurationSec:        durationSec(14, 25),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "830",
			OriginAirport:      api.NewAirportSpec(0, "SYD"),
			DestinationAirport: api.NewAirportSpec(0, "SFO"),
			Fleet:              api.NewFleetSpec(0, "B789"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "11:30",
			DurationSec:        durationSec(14, 35),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "5703",
			OriginAirport:      api.NewAirportSpec(0, "SFO"),
			DestinationAirport: api.NewAirportSpec(0, "MSP"),
			Fleet:              api.NewFleetSpec(0, "A320"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "08:30",
			DurationSec:        durationSec(4, 15),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "5297",
			OriginAirport:      api.NewAirportSpec(0, "MSP"),
			DestinationAirport: api.NewAirportSpec(0, "SFO"),
			Fleet:              api.NewFleetSpec(0, "A320"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "15:30",
			DurationSec:        durationSec(4, 45),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "14",
			OriginAirport:      api.NewAirportSpec(0, "EWR"),
			DestinationAirport: api.NewAirportSpec(0, "LHR"),
			Fleet:              api.NewFleetSpec(0, "B78X"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "18:30",
			DurationSec:        durationSec(6, 15),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "110",
			OriginAirport:      api.NewAirportSpec(0, "EWR"),
			DestinationAirport: api.NewAirportSpec(0, "LHR"),
			Fleet:              api.NewFleetSpec(0, "B78X"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "22:30",
			DurationSec:        durationSec(6, 15),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "59",
			OriginAirport:      api.NewAirportSpec(0, "FRA"),
			DestinationAirport: api.NewAirportSpec(0, "SFO"),
			Fleet:              api.NewFleetSpec(0, "B77W"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "10:15",
			DurationSec:        durationSec(12, 30),
		},
		{
			Airline:            api.NewAirlineSpec(0, "KL"),
			Number:             "1823",
			OriginAirport:      api.NewAirportSpec(0, "AMS"),
			DestinationAirport: api.NewAirportSpec(0, "FRA"),
			Fleet:              api.NewFleetSpec(0, "E190"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "07:45",
			DurationSec:        durationSec(0, 53),
		},
		{
			Airline:            api.NewAirlineSpec(0, "BA"),
			Number:             "430",
			OriginAirport:      api.NewAirportSpec(0, "LHR"),
			DestinationAirport: api.NewAirportSpec(0, "AMS"),
			Fleet:              api.NewFleetSpec(0, "A320"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "06:45",
			DurationSec:        durationSec(0, 59),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "344",
			OriginAirport:      api.NewAirportSpec(0, "SFO"),
			DestinationAirport: api.NewAirportSpec(0, "DFW"),
			Fleet:              api.NewFleetSpec(0, "A320"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "09:30",
			DurationSec:        durationSec(4, 45),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "271",
			OriginAirport:      api.NewAirportSpec(0, "DFW"),
			DestinationAirport: api.NewAirportSpec(0, "SFO"),
			Fleet:              api.NewFleetSpec(0, "A320"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "16:30",
			DurationSec:        durationSec(4, 59),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "1954",
			OriginAirport:      api.NewAirportSpec(0, "SFO"),
			DestinationAirport: api.NewAirportSpec(0, "DCA"),
			Fleet:              api.NewFleetSpec(0, "B738"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "13:30",
			DurationSec:        durationSec(5, 15),
		},
		{
			Airline:            api.NewAirlineSpec(0, "UA"),
			Number:             "395",
			OriginAirport:      api.NewAirportSpec(0, "DCA"),
			DestinationAirport: api.NewAirportSpec(0, "SFO"),
			Fleet:              api.NewFleetSpec(0, "B738"),
			StartDate:          "2025-01-25",
			EndDate:            "2025-02-07",
			DaysOfWeek:         []int{0, 1, 2, 3, 4, 5, 6},
			DepartureTime:      "07:30",
			DurationSec:        durationSec(5, 45),
		},
	}
	for _, f := range flightSchedules {
		f.Published = ptrTo(true)
	}
	log.Println("Creating flight schedules...")
	for _, f := range flightSchedules {
		if _, err := handler.CreateFlightSchedule(ctx, api.CreateFlightScheduleRequestObject{Body: f}); err != nil {
			return fmt.Errorf("inserting flight schedule: %w", err)
		}
	}

	log.Println("Creating passengers...")
	passengerNames := []string{
		"John Doe", "Jane Doe", "Bob Smith", "John Smith", "Alice Zhao", "Maria Garcia", "James Johnson", "Sarah Wilson", "Michael Chen", "Emily Brown", "David Kim", "Lisa Patel", "Carlos Rodriguez", "Emma Davis", "Mohammed Ahmed", "Sofia Martinez", "William Lee",
		"Olivia Taylor", "Daniel Jackson", "Isabella Lopez", "Alexander Wong", "Ava Thompson", "Lucas Nguyen", "Mia Anderson", "Ethan Kumar", "Sophia White", "Ryan O'Connor", "Grace Williams", "Nathan Cohen", "Victoria Singh",
		"Liam Johnson", "Avery Thompson", "Elijah Martinez", "Scarlett Davis", "William Wang", "Chloe Anderson", "Noah Hernandez", "Camila Rodriguez", "Oliver Garcia", "Evelyn Hernandez", "Lucas Wilson", "Mila King", "James Brown", "Zoe Lee", "Benjamin Lewis", "Aria Young", "Henry Miller",
		// "Penelope Davis", "Joseph Thompson", "Grace Davis", "Nathan Garcia", "Aria Rodriguez", "Mia Wilson", "Camila Anderson", "Ethan Martinez", "Olivia White",
		// "Logan Walker", "Abigail Harris", "Samuel Green", "Avery Turner", "Joseph Hill", "Mila Foster", "Henry Campbell", "Sofia Reyes", "Carter Rivera", "Evelyn Cooper",
		// "Thomas Mitchell", "Grace Turner", "Elijah Bailey", "Zoe Bailey", "Ella Lee", "Aiden Davis", "Avery Johnson", "Aubrey Wilson", "Cadence Perez", "Hannah Morris",
	}
	passengerIDs := make([]int, len(passengerNames))
	for i, name := range passengerNames {
		resp, err := handler.CreatePassenger(ctx, api.CreatePassengerRequestObject{Body: &api.CreatePassengerJSONRequestBody{
			Name: name,
		}})
		if err != nil {
			return fmt.Errorf("inserting passenger: %w", err)
		}
		passengerIDs[i] = resp.(api.CreatePassenger201JSONResponse).Id
	}

	log.Println("Creating itineraries...")
	flightInstances, err := handler.ListFlightInstances(ctx, api.ListFlightInstancesRequestObject{})
	if err != nil {
		return err
	}
	itinsCreated := 0
	for _, f := range flightInstances.(api.ListFlightInstances200JSONResponse)[:10] {
		for _, passengerID := range passengerIDs {
			_, err := handler.CreateItinerary(ctx, api.CreateItineraryRequestObject{
				Body: &api.CreateItineraryJSONRequestBody{
					FlightInstanceIDs: []int{f.Id},
					PassengerIDs:      []int{passengerID},
				},
			})
			if err != nil {
				return fmt.Errorf("inserting itinerary: %w", err)
			}
			itinsCreated++
			if itinsCreated > 0 && itinsCreated%100 == 0 {
				log.Printf("- created %d itineraries", itinsCreated)
			}
		}
	}

	log.Println("Creating seat assignments...")
	itineraries, err := handler.ListItineraries(ctx, api.ListItinerariesRequestObject{})
	if err != nil {
		return err
	}
	seatAssignmentsCreated := 0
	for _, itin := range itineraries.(api.ListItineraries200JSONResponse) {
		seatRow := rand.Intn(30) + 1
		seatLetter := 'A' + rand.Intn(10)
		randomSeat := fmt.Sprintf("%d%c", seatRow, seatLetter)
		_, err := handler.CreateSeatAssignment(ctx, api.CreateSeatAssignmentRequestObject{
			FlightInstanceID: itin.Flights[0].Id,
			Body: &api.CreateSeatAssignmentJSONRequestBody{
				ItineraryID: itin.Id,
				PassengerID: itin.Passengers[0].Id,
				Seat:        randomSeat,
			},
		})
		if err != nil {
			// TODO!(sqs)
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				continue
			}
			return err
		}
		seatAssignmentsCreated++
		if seatAssignmentsCreated > 0 && seatAssignmentsCreated%100 == 0 {
			log.Printf("- created %d seat assignments", seatAssignmentsCreated)
		}
	}

	return nil
}
