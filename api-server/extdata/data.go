//go:build !gen
// +build !gen

package extdata

import (
	"bytes"
	"encoding/gob"

	_ "embed"
)

// TODO!(sqs): make this not run each time
//OFF go:generate go run -tags gen github.com/stellora/airline/api-server/extdata/get-extdata -format=gob -out=airports.data.gob

//go:embed airports.data.gob
var airportsData []byte

var Airports AirportsDataset

func LoadAirports() error {
	return gob.NewDecoder(bytes.NewReader(airportsData)).Decode(&Airports)
}
