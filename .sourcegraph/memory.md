# Build and Test Commands

## Web (Frontend)
- **Install dependencies:** `pnpm -C web install`
- **Development server:** `pnpm -C web dev` (then visit http://localhost:5174)
  - Assume the user is already running the dev server at http://localhost:5174. For any changes to the web app, ALWAYS navigate to the page on http://localhost:5174 in Puppeteer and confirm the changes.
- **Build:** `pnpm -C web build`
- **Lint:** `pnpm -C web lint`
- **Type check:** `pnpm -C web check`
- **Format:** `pnpm -C web format`
- **Run all tests:** `pnpm -C web test`
- **Run unit tests:** `pnpm -C web test:unit`
- **Run single unit test:** `pnpm -C web test:unit -t "test name"`
- **Run E2E tests:** `pnpm -C web test:e2e`

## API Server (Backend)
- **Run server:** `cd api-server && go run .`
- **Run all tests:** `cd api-server && go test ./...`
- **Run specific test:** `cd api-server && go test -v ./... -run TestName`
- **Run specific test file:** `cd api-server && go test -v ./file_test.go`

## Code Style
- **Frontend:** TypeScript/Svelte with tabs, single quotes, no semi-colons, 100 char width
- **Backend:** Go with standard formatting (`gofmt`)
- **Test fixtures:** Airlines (XX, YY, ZZ), Airports (AAA, BBB, CCC), Aircraft types use real codes
- **Rendering:** Flight numbers and IATA codes in monospace (`font-mono` class)
- **Error handling:** Go uses explicit error returns, Frontend uses async/await with try/catch
- **Conventions:** Svelte components use PascalCase, utility functions use camelCase
