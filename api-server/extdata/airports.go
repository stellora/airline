package extdata

import (
	_ "embed"
)

type Continent string

var Continents = map[Continent]string{
	"AF": "Africa",
	"AN": "Antarctica",
	"AS": "Asia",
	"EU": "Europe",
	"NA": "North America",
	"OC": "Oceania",
	"SA": "South America",
}

type ISOCountry string // 2-character ISO 3166:1-alpha2 code for a country

// Country is an entry in https://ourairports.com/help/data-dictionary.html#countries.
type Country struct {
	ID            int
	Code          ISOCountry // 2-character ISO 3166:1-alpha2 code for the country
	Name          string
	Continent     Continent
	WikipediaLink string   `json:",omitempty"`
	Keywords      []string `json:",omitempty"`
}

type ISORegion string // ISOCountry + region identifier

// Region is an entry in https://ourairports.com/help/data-dictionary.html#regions.
type Region struct {
	ID            int
	Code          ISORegion
	LocalCode     string
	Name          string
	Continent     Continent
	ISOCountry    ISOCountry // 2-character ISO 3166:1-alpha2 code for the country (maps to (Country).Code)
	WikipediaLink string     `json:",omitempty"`
	Keywords      []string   `json:",omitempty"`
}

// Airport is an entry in https://ourairports.com/help/data-dictionary.html#airports.
type Airport struct {
	ID               int
	Ident            string
	Type             string
	Name             string
	LatitudeDeg      float64
	LongitudeDeg     float64
	ElevationFt      *int `json:",omitempty"`
	Continent        Continent
	ISOCountry       ISOCountry
	ISORegion        ISORegion
	Municipality     string
	ScheduledService bool
	GPSCode          string
	IATACode         string
	LocalCode        string   `json:",omitempty"`
	HomeLink         string   `json:",omitempty"`
	WikipediaLink    string   `json:",omitempty"`
	Keywords         []string `json:",omitempty"`
}

type AirportsDataset struct {
	Countries map[ISOCountry]Country
	Regions   map[ISORegion]Region
	Airports  []Airport
}
