# Airline (demo app)

## Usage

```shell
# Install dependencies and verify tests pass.
pnpm -C web install
pnpm -C web test

# Start a dev server.
pnpm -C web dev
```

## Development

- Codegen: `pnpm -C web gen && (cd api-server && go generate ./...)`