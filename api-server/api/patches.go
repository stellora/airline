package api

// MarshalText implements encoding.TextMarshaler interface for AirlineSpec
func (a AirlineSpec) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface for AirlineSpec
func (a *AirlineSpec) UnmarshalText(text []byte) error {
	str := string(text)
	val := AirlineSpec(str)
	*a = val
	return nil
}
