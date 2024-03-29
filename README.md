# WeatherApp using golang and ginkgo

Designed to conceptually understand Go

- Structs, Pointers and Constants
- Interfaces
- Methods and Functions
- Package Management
- http client
- JSON Encoding/Decoding (Marshalling / Unmarshalling )

# Work In Progress


## Setup and WeatherReport Struct

- [x] Set up development environment with Ginkgo and Gomega testing frameworks.
- [x] Create a new project structure with necessary directories (e.g., src, test).
- [x] Initialize a Git repository and commit initial project setup.
- [x] Define the `WeatherReport` struct with fields for weather data (e.g., temperature, humidity, wind speed).
- [x] Implement basic methods for initializing and accessing weather data within the struct.

## Struct Validator and Custom Defined Rule(using go playground)

- [x] Implement a struct validator to validate the `WeatherReport` struct against defined rules (e.g., temperature range, humidity range) using Ginkgo and Gomega.
- [x] Define custom validation rules for specific fields (e.g., wind speed must be non-negative).
- [x] Write Ginkgo tests using Gomega assertions for the struct validator to ensure correct validation behavior.
- [x] Test different scenarios including valid data and data violating custom-defined rules.

## API Integration

- [x] Integrate the [Open Meteo API](https://api.open-meteo.com/v1/forecast) to collect weather report data for a specific location.
- [x] Implement methods to make API requests and retrieve weather data in `WeatherReport` struct format.
- [x] Write Ginkgo tests using Gomega assertions to validate the API integration.
- [x] Test scenarios covering successful API requests and handling of errors or invalid responses.

## Test Coverage, Refinement, and Documentation

- [ ] Review and refine codebase for clarity, consistency, and adherence to best practices.
- [ ] Improve test coverage by adding additional Ginkgo tests with Gomega assertions to cover edge cases and error handling scenarios.
- [ ] Conduct thorough testing of the entire codebase using Ginkgo and Gomega, including unit tests, integration tests, and end-to-end tests.
- [ ] Ensure that all Ginkgo tests pass successfully and address any failing tests by fixing code issues.

## Beyond Go

- [ ] Github Actions
- [ ] Terraform
- [ ] Helm

