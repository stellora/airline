# Airline (demo app)

Airline is a demonstration application that simulates an airline reservation and flight management system. It provides functionality for:

- Managing flight schedules and routes between airports
- Tracking airlines, aircraft, and fleet information
- Handling passenger itineraries and bookings
- Visualizing flight routes and airport connections

The application consists of a Go-based API server and a TypeScript/web frontend.

## Usage

```shell
# Install dependencies and verify tests pass.
pnpm -C web install
pnpm -C web test
(cd api-server && go test ./...)

# Start a dev server by running both commands in separate windows:
(cd api-server && go run .)
pnpm -C web dev
```

## Development

See `mise run` for development commands.

## Data Sources

The application uses real-world data sources:

- [Airports data](https://ourairports.com/data/) from OurAirports - provides accurate airport information including IATA codes and locations
- Sample flight schedules and routes based on real airline operations
- Simulated booking and passenger data for demonstration purposes

## Style guide

- Offer both light and dark themes.
- Always render flight numbers (e.g., `UA123`) and IATA airport codes (e.g., `SFO`) in monospace (`font-mono` class).
- DateTime means TODO!(sqs), time and date mean 
- localtime means a tz-less time like "6pm", localdate means a tz-less date like "January 25, 2025"

## Notes

- See https://standards.atlassian.net/wiki/spaces/AIDM/pages/607649825/IATA+Open+Air+JSON+schema+library for a JSON Schema.
- Public stats: https://transtats.bts.gov/databases.asp?Z1qr_VQ=E&Z1qr_Qr5p=N8vn6v10&f7owrp6_VQF=D.