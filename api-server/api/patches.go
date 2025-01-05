package api

import (
	"encoding"
	"regexp"
	"strconv"
)

var intString = regexp.MustCompile(`^\d+$`)

// isIntString returns true if str is a string of one or more digits. Unlike strconv.Atoi or
// strconv.ParseInt, it does not allow leading '-' or '+' characters.
func isIntString(str string) bool {
	return intString.MatchString(str)
}

func (a *AircraftSpec) UnmarshalText(text []byte) error {
	*a = aircraftSpecFromPathArg(string(text))
	return nil
}

var _ encoding.TextUnmarshaler = (*AircraftSpec)(nil)

func aircraftSpecFromPathArg(arg string) AircraftSpec {
	if isIntString(arg) {
		id, _ := strconv.Atoi(arg)
		return NewAircraftSpec(id, "")
	}
	return NewAircraftSpec(0, arg)
}

func NewAircraftSpec(id int, registration string) AircraftSpec {
	var spec AircraftSpec
	if id != 0 {
		spec.FromAircraftID(id)
	} else {
		spec.FromAircraftRegistration(registration)
	}
	return spec
}

func (a *AirportSpec) UnmarshalText(text []byte) error {
	*a = airportSpecFromPathArg(string(text))
	return nil
}

var _ encoding.TextUnmarshaler = (*AirportSpec)(nil)

func airportSpecFromPathArg(arg string) AirportSpec {
	if isIntString(arg) {
		id, _ := strconv.Atoi(arg)
		return NewAirportSpec(id, "")
	}
	return NewAirportSpec(0, arg)
}

func NewAirportSpec(id int, iataCode string) AirportSpec {
	var spec AirportSpec
	if id != 0 {
		spec.FromAirportID(id)
	} else {
		spec.FromAirportIATACode(iataCode)
	}
	return spec
}

func (a *AirlineSpec) UnmarshalText(text []byte) error {
	*a = airlineSpecFromPathArg(string(text))
	return nil
}

var _ encoding.TextUnmarshaler = (*AirlineSpec)(nil)

func airlineSpecFromPathArg(arg string) AirlineSpec {
	if isIntString(arg) {
		id, _ := strconv.Atoi(arg)
		return NewAirlineSpec(id, "")
	}
	return NewAirlineSpec(0, arg)
}

func NewAirlineSpec(id int, iataCode string) AirlineSpec {
	var spec AirlineSpec
	if id != 0 {
		spec.FromAirlineID(id)
	} else {
		spec.FromAirlineIATACode(iataCode)
	}
	return spec
}

func (a *ItinerarySpec) UnmarshalText(text []byte) error {
	*a = itinerarySpecFromPathArg(string(text))
	return nil
}

var _ encoding.TextUnmarshaler = (*ItinerarySpec)(nil)

func itinerarySpecFromPathArg(arg string) ItinerarySpec {
	if isIntString(arg) {
		id, _ := strconv.Atoi(arg)
		return NewItinerarySpec(id, "")
	}
	return NewItinerarySpec(0, arg)
}

func NewItinerarySpec(id int, recordLocator string) ItinerarySpec {
	var spec ItinerarySpec
	if id != 0 {
		spec.FromItinerarySpec0(id)
	} else {
		spec.FromRecordLocator(recordLocator)
	}
	return spec
}
