# Go Service (Gin)

This is a Go microservice using the Gin framework for the Best Egg interview challenge requirements.

## Project Setup

1. Ensure Go 1.20+ is installed.
2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Running the Mock Server

The service depends on the provided WireMock mock server.

1. From the `interview-challenge-2025` directory, run:
   ```sh
   docker compose up
   ```
2. WireMock will be available at [http://localhost:8080](http://localhost:8080)

## Running the Go Service

From the `go-service` directory:

```sh
go run main.go
```

The service will be available at [http://localhost:3000](http://localhost:3000)

Swagger/OpenAPI docs: [http://localhost:3000/docs/index.html](http://localhost:3000/docs/index.html)

## API Endpoints

### `GET /packages`
- Returns a list of all known packages.
- Supports filtering by status:
  - Example: `/packages?status=Delivered`
- Supports sorting by `eta` or `last_updated`:
  - Example: `/packages?sort=eta`
  - Example: `/packages?status=In%20Transit&sort=last_updated`

### `GET /packages/{tracking_id}`
- Returns detailed tracking info for a package.

### `GET /carriers`
- Returns a list of supported carriers.

## Running Tests

The code is structured for easy unit testing. Handler dependencies (such as data-fetching functions) are assigned to variables, so they can be mocked in tests.

To run all tests:
```sh
go test
```

You can add more tests in `routes_test.go` using the same mocking approach. 