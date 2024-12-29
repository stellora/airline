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