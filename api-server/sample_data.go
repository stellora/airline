package main

import "context"

func insertSampleData(ctx context.Context, handler *Handler) error {
	airports := []string{"SFO", "SIN", "NRT", "HND", "EWR", "LAX", "DEN", "ORD", "AMS", "LHR", "HKG", "SYD", "HNL", "LIH", "SJC", "STS", "SEA", "FRA", "MUC", "DXB", "TLV", "IST", "DOH", "DEL", "BOM", "KIX", "MEL", "JNB", "CPT", "EZE", "MEX", "CUN", "MSP", "IAD", "DCA", "DFW", "MRS", "PVR", "BOS", "FCO", "SCL", "CDG"}
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
	}
	if _, err := insertAirlines(ctx, handler, airlines); err != nil {
		return err
	}

	flightTitles := []string{
		"UA1 SFO-SIN", "UA2 SIN-SFO",
		"SQ33 SFO-SIN",
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
