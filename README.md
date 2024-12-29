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