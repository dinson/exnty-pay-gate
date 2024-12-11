# Payment Gateway Integration Assessment

This project handles deposit and withdrawal transactions using 3rd party payment gateways.

### API Reference

Visit https://exinity-payment-gateway.readme.io/

OpenAPI spec can be found in the directory `./docs/api.yaml`

### Setup instructions

1. **Clone the Repository:**
    Clone the repository to your local machine:

    ```bash
    git clone https://github.com/dinson/exnty-pay-gate
    cd exnty-pay-gate
    ```

2. **Setup Docker:**
    Docker is configured to run PostgreSQL, Kafka, and Redis. Use the following command to start all the services:

    ```bash
    docker-compose up -d
    ```

    This will start:
    - PostgreSQL on port `5432`
    - Kafka on ports `9092` and `9093`
    - Redis on port `6379`
    - Application on port `8080`

3. **Database Migration:**
    The migration file `db/init.sql` is already provided. Once the Docker services are up and running, the database will be initialized automatically, and the tables will be created.


### Architecture

 - The project follows "Layered" architecture, ensuring proper separation of concerns, modularity, unit-testability and maintainability.

### Unit testing

 - In project root, run `go test ./...` to trigger unit testing of all packages.
 - The project uses `github.com/stretchr/testify` package to assert test results.
 - Uses `vektra/mockery` to generate mocks. `https://github.com/vektra/mockery`
 - Install mockery by running `brew install mockery`
 - Mocks for a package can be generated by running `mockery --name=Interface`

### Important packages
- **config**
  - Stores the configuration values including credentials and other values required in the project.
- **client**
  - Create and store all the clients required in the project.
- **internal / api / handler**
  - Handle all API requests.
  - Defines `protected` and `public` endpoints.
  - `public` endpoints are essentially webhook callbacks from gateways.
- **internal / middleware**
  - Defines middleware functions.
  - Verifies the user, extract and store user's country to request context.
- **internal / services**
  - Business logic for all the operations.
  - Every package is exposed through an interface, abstracting its implementation.
- **paymentprovider**
  - Encapsulates all 3rd party payment provider libraries, providing a unified facade interface to the outside world.
  - Extensible to add new payment provider packages without modifying the existing code.
- **db**
  - Exposes interface for db package hiding the method implementations.
- **constant**
  - Defines all constant values required in the program.
- **context**
  - Defines utility functions that work around the values stored in the context.Context.
- **enum**
  - Defines all enumerated values used in the program.
- **errors**
  - Define all the errors used in the program.
- **docs**
  - Documentation for workflows and OpenAPI spec.

## Gateway Selection and transactions

- Each gateway is mapped to the available countries in the `gateway_priority` table.
- Every row will hold the priority value for each gateway for a country.
- During deposit or withdrawal transaction requests:
  - User's country is retrieved in the middleware and stored in the context.
  - When retrieving the gateways for a user's country, the DB method `ListCountryGatewaysByPriority()` will return the gateways in the ascending order of their priority.
  - the available gateways will be invoked one after the other in the order of their priority. If any of the gateway returns successful response, the rest of the gateway list is ignored.
  - If there are no gateways configured for the user's country, the handler will return an error.