# Airline (demo app)

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

## Data

- [Airports data](https://ourairports.com/data/) from OurAirports

## Style guide

- Offer both light and dark themes.
- Always render flight numbers (e.g., `UA123`) and IATA airport codes (e.g., `SFO`) in monospace (`font-mono` class).
- DateTime means TODO!(sqs), time and date mean 
- localtime means a tz-less time like "6pm", localdate means a tz-less date like "January 25, 2025"

## Notes

- See https://standards.atlassian.net/wiki/spaces/AIDM/pages/607649825/IATA+Open+Air+JSON+schema+library for a JSON Schema.
- Public stats: https://transtats.bts.gov/databases.asp?Z1qr_VQ=E&Z1qr_Qr5p=N8vn6v10&f7owrp6_VQF=D.