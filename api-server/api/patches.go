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

// UnmarshalText implements encoding.TextUnmarshaler interface for AirportSpec
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
