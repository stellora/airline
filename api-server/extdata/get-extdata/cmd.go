package main

import (
	"context"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/stellora/airline/api-server/extdata"
)

var (
	encodeFormat = flag.String("format", "gob", "encode to \"gob\" or \"json\" format")
	outFile      = flag.String("out", "airports.data.gob", "output file name")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	countries, err := readCountries(ctx)
	if err != nil {
		log.Fatal(err)
	}

	airports, err := readAirports(ctx, func(airport *extdata.Airport) bool {
		return airport.ScheduledService || airport.Type == "large_airport" || airport.Type == "medium_airport"
	})
	if err != nil {
		log.Fatal(err)
	}

	seenRegions := map[extdata.ISORegion]struct{}{}
	for _, airport := range airports {
		seenRegions[airport.ISORegion] = struct{}{}
	}

	regions, err := readRegions(ctx, func(region *extdata.Region) bool {
		if region.ISOCountry == "US" {
			return true
		}
		_, seen := seenRegions[region.Code]
		return seen
	})
	if err != nil {
		log.Fatal(err)
	}

	dataset := &extdata.AirportsDataset{
		Countries: countries,
		Regions:   regions,
		Airports:  airports,
	}

	out, err := os.Create(*outFile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	switch *encodeFormat {
	case "gob":
		if err := gob.NewEncoder(out).Encode(dataset); err != nil {
			log.Fatal(err)
		}
	case "json":
		if err := json.NewEncoder(out).Encode(dataset); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown encode format %q", *encodeFormat)
	}
}

func readCountries(ctx context.Context) (map[extdata.ISOCountry]extdata.Country, error) {
	file, err := download(ctx, "https://ourairports.com/data/countries.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	countries := make(map[extdata.ISOCountry]extdata.Country, 250)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6
	reader.ReuseRecord = true

	// Skip header row
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("reading countries CSV header: %w", err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("reading countries CSV: %w", err)
		}

		code := extdata.ISOCountry(record[1])
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("parsing country ID: %w", err)
		}
		countries[code] = extdata.Country{
			ID:            id,
			Code:          code,
			Name:          record[2],
			Continent:     extdata.Continent(record[3]),
			WikipediaLink: record[4],
			Keywords:      splitKeywords(record[5]),
		}
	}

	return countries, nil
}

func readRegions(ctx context.Context, filter func(region *extdata.Region) bool) (map[extdata.ISORegion]extdata.Region, error) {
	file, err := download(ctx, "https://ourairports.com/data/regions.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	regions := make(map[extdata.ISORegion]extdata.Region, 4000)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 8
	reader.ReuseRecord = true

	// Skip header row
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("reading regions CSV header: %w", err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("reading regions CSV: %w", err)
		}

		code := extdata.ISORegion(record[1])
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("parsing region ID: %w", err)
		}
		region := extdata.Region{
			ID:            id,
			Code:          code,
			LocalCode:     record[2],
			Name:          record[3],
			Continent:     extdata.Continent(record[4]),
			ISOCountry:    extdata.ISOCountry(record[5]),
			WikipediaLink: record[6],
			Keywords:      splitKeywords(record[7]),
		}
		if filter == nil || filter(&region) {
			regions[code] = region
		}
	}

	return regions, nil
}

func readAirports(ctx context.Context, filter func(airport *extdata.Airport) bool) ([]extdata.Airport, error) {
	file, err := download(ctx, "https://ourairports.com/data/airports.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	airports := make([]extdata.Airport, 0, 10000)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 18
	reader.ReuseRecord = true

	// Skip header row
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("reading airports CSV header: %w", err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("reading airports CSV: %w", err)
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("parsing airport ID: %w", err)
		}
		latitudeDeg, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return nil, fmt.Errorf("parsing latitude: %w", err)
		}
		longitudeDeg, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			return nil, fmt.Errorf("parsing longitude: %w", err)
		}
		var elevationFtPtr *int
		if record[6] != "" {
			elevationFt, err := strconv.Atoi(record[6])
			if err != nil {
				return nil, fmt.Errorf("parsing elevation: %w", err)
			}
			elevationFtPtr = &elevationFt
		}

		airport := extdata.Airport{
			ID:               id,
			Ident:            record[1],
			Type:             record[2],
			Name:             record[3],
			LatitudeDeg:      latitudeDeg,
			LongitudeDeg:     longitudeDeg,
			ElevationFt:      elevationFtPtr,
			Continent:        extdata.Continent(record[7]),
			ISOCountry:       extdata.ISOCountry(record[8]),
			ISORegion:        extdata.ISORegion(record[9]),
			Municipality:     record[10],
			ScheduledService: record[11] == "yes",
			GPSCode:          record[12],
			IATACode:         record[13],
			LocalCode:        record[14],
			HomeLink:         record[15],
			WikipediaLink:    record[16],
			Keywords:         splitKeywords(record[17]),
		}
		if filter == nil || filter(&airport) {
			airports = append(airports, airport)
		}
	}

	return airports, nil
}

func splitKeywords(keywords string) []string {
	if keywords == "" {
		return nil
	}
	return strings.Split(keywords, ",")
}

func download(ctx context.Context, url string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("get %s: unexpected status %s", req.URL, resp.Status)
	}
	return resp.Body, nil
}
