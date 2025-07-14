# Best Egg Engineering Code Challenge

Your task is to build a FastAPI-based microservice that integrates with a mocked third-party package tracking API.

## Prep
- Create a repository using your personal GitHub account
- Copy the provided docker-compose file and wiremock directory from this repository into your own
- Run the mock server using docker compose
    - WireMock will be available at `http://localhost:8080`

## Your Microservice Should Include

- `GET /packages/{tracking_id}` – Fetch detailed tracking info using the mock API
- `GET /packages` – Return a paginated list of all known packages
    - Supports filtering by status (e.g., `?status=Delivered`)
    - Supports sorting by `eta` or `last_updated` (e.g., `?sort=eta`)
    - Supports pagination with `page` and `limit` query parameters (e.g., `?page=2&limit=5`)
    - The response includes:
      ```json
      {
        "packages": [ ... ],
        "page": 1,
        "limit": 10,
        "total": 42
      }
      ```
- `GET /carriers` – Return a list of supported carriers from the mock API

## Requirements

- Use FastAPI with OpenAPI docs enabled at `/docs`
- Use Pydantic models for request/response validation
- Include unit tests with `pytest`

## Stretch Goals

- Join data from `/tracking/{id}` and `/locations/{city}` to enrich the tracking response with city metadata
- Implement integration testing
- Retry failed requests (e.g., 500 responses) with backoff
- Add pagination or result limiting
- Add something not listed above! A differentiator that will make your project stand out

## Notes

- You may use any tools you'd use day-to-day (e.g., ChatGPT, Copilot)
- You may add additional data to the mocked API to demonstrate your project's capabilities, but preserve the core endpoints and structure
- You'll walk through your implementation with us during the interview
- You may be asked to implement additional enhancements during the inteview, be prepared to share your screen and work on your code
