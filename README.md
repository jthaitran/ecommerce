# WooCommerce Product Management

This project provides a Go-based solution for managing products in WooCommerce using the WooCommerce API (Version 3). It follows a clean architectural design and provides a modular structure that allows for easy maintenance, testing, and scalability.

## Prerequisites

- Go 1.16 or higher installed on your machine.
- WooCommerce store with API access enabled.
- WooCommerce consumer key and secret.

## Project Structure

The project follows a modular structure that promotes separation of concerns and maintainability. Here's an overview of the folder structure:

- main.go: This is the entry point of your application.
- internal/app/product/: This directory contains the application-specific code for managing products.
- handler.go: Defines the HTTP handlers for product-related endpoints.
- service.go: Implements the business logic for product management.
- repository.go: Provides an interface and its implementation for accessing the product data.
- internal/domain/: Contains the domain models and interfaces that define the core concepts of your application.
- product.go: Defines the product entity and related interfaces.
- internal/integrations/woocommerce/: Includes the integration-specific code for interacting with the WooCommerce API.
- client.go: Implements the client for making requests to the WooCommerce API.
- product_repository.go: Implements the product repository interface using the WooCommerce API client.
- test/app/product/: Contains the unit tests for the application-level code related to product management.
- handler_test.go: Tests the HTTP handlers for product-related endpoints.
- service_test.go: Tests the business logic for product management.
- repository_test.go: Tests the product repository implementation.
- test/domain/: Includes the unit tests for the domain models and interfaces.
- product_test.go: Tests the behavior of the product entity and related interfaces.
- test/integrations/woocommerce/: Contains the integration tests for the WooCommerce integration.
- product_repository_test.go: Tests the product repository implementation using the WooCommerce API client.


## Installation

1. Clone the repository:
```bash
git clone https://github.com/jthaitran/ecommerce.git
```
2. Change into the project directory:
```bash
cd woocommerce-product-management
```
3. Install the dependencies:
```bash
go mod download
```

## Configuration

1. Open the `config.go` file located in the `config` directory.

2. Update the following configuration options with your WooCommerce store details:

```go
const (
    BaseURL       = "https://your-store-url.com"  // Replace with your WooCommerce store URL
    ConsumerKey   = "your-consumer-key"           // Replace with your WooCommerce consumer key
    ConsumerSecret = "your-consumer-secret"       // Replace with your WooCommerce consumer secret
)
```
Usage
To run the application, use the following command:
```cmd
go run cmd/main.go
```

You can also build the application into an executable file:
```cmd
go build -o woocommerce-product-management cmd/main.go
```
## Architectural Design
The project follows the Clean Architecture design principles to promote separation of concerns and maintainability. Here's a brief overview of the architecture:

- The domain package contains the core business logic, including entities, models, and interfaces.
- The usecase package implements the application-specific use cases and interacts with the domain layer.
- The infrastructure package provides implementations for interacting with external dependencies, such as the WooCommerce API.
- The cmd package serves as the entry point for the application and provides a command-line interface.
```cmd
- main.go
- internal/
  - app/
    - product/
      - handler.go
      - service.go
      - repository.go
  - domain/
    - product.go
  - integrations
    - woocommerce/
      - client.go
      - product_repository.go
- test/
  - app/
    - product/
      - handler_test.go
      - service_test.go
      - repository_test.go
  - domain/
    - product_test.go
  - integrations/
    - woocommerce/
      - product_repository_test.go
```


## Unit Testing
The project emphasizes unit testing to ensure the correctness of the implemented functionality. Unit tests are located in the respective packages and can be run using the following command:

```bash
go test ./...
```
Feel free to add more unit tests to the project to cover additional functionality and edge cases.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request. - Tran Quang Thai

## License
This project is licensed under the MIT License.
