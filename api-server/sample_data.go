package main

import (
	"context"

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
		{Registration: "N27251", AircraftType: "B38M", Airline: ua},
		{Registration: "N27252", AircraftType: "B38M", Airline: ua},
		{Registration: "N27253", AircraftType: "B38M", Airline: ua},
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
	}
	for _, a := range aircraft {
		if _, err := handler.CreateAircraft(ctx, api.CreateAircraftRequestObject{Body: &a}); err != nil {
			return err
		}
	}

	flightTitles := []string{
		"UA1 SFO-SIN", "UA2 SIN-SFO",
		"SQ33 SFO-SIN", "SQ34 SIN-SFO",
		"SQ31 SFO-SIN", "SQ32 SIN-SFO",
		"SQ345 ZRH-SIN", "SQ346 SIN-ZRH",
		"UA2168 SFO-EWR", "UA1054 SFO-EWR", "UA2460 SFO-EWR",
		"UA2855 EWR-SFO", "UA598 EWR-SFO", "UA1579 EWR-SFO",
		"UA1684 SFO-LIH", "UA1563 LIH-SFO", "UA1111 SFO-LIH",
		"UA863 SFO-SYD", "UA830 SYD-SFO",
		"UA5703 SFO-MSP", "UA5297 MSP-SFO",
		"UA14 EWR-LHR", "UA110 EWR-LHR", "UA59 FRA-SFO",
		"KL1823 AMS-FRA", "BA430 LHR-AMS",
		"UA344 SFO-DFW", "UA271 DFW-SFO",
		"UA1954 SFO-DCA", "UA395 DCA-SFO",
		"UA968 AMS-SFO", "UA969 AMS-SFO",
		"LH1091 MRS-FRA",
		"UA374 SFO-IAD",
		"UA926 SFO-FRA", "UA927 FRA-SFO",
		"LH458 MUC-SFO",
		"UA40 EWR-FCO",
		"UA194 SFO-MUC",
		"DL2649 MSP-BOS",
		"UA1212 SFO-PVR", "UA1243 PVR-SFO",
	}
	if _, err := insertFlightSchedules(ctx, handler, flightTitles...); err != nil {
		return err
	}

	return nil
}
