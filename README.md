# Lido events

## Overview

**Description**: This application is a Go-based service designed with hexagonal architecture to manage and interact with Lido Operators. It includes the following features:

- Subscribing to the Ethereum events based on the operator IDs loaded in the app
- Scanning Ethereum events based on the operator IDs loaded in the app:
  - `ExitRequested`: event emitted by the CSFeeDistrbutor SC that indicates whenever a validator requires to exit the CSM pool.
  - `DistributedLogUpdated`: event emitted by the VEBO SC that indicates whenever a new log is distributed. This log contains the performance data of the validators as well as the treshold.
- Executing validator exits in background when there is a exit request event requested for a validator of the operator ID loaded in the app
- Fetching and parsing IPFS log CIDs in background to get the validator performance data
- Sending notifications to user (e.g., via Telegram)
- Exposing an API to update configurations and retrieve validator performance data

## Architecture

**Hexagonal Architecture**: The app follows hexagonal architecture principles, decoupling core business logic from external dependencies, enhancing modularity and maintainability. This design enables easy replacement or modification of adapters (e.g., Ethereum clients, IPFS, notification systems).

## Features

- **Event Subscriptions**: The app subscribes to a list of Ethereum events.
- **Event Scanning**: It regularly scans specific Ethereum events.
- **Validator Exits**: Allows for executing validator exits.
- **IPFS Data Parsing**: Fetches and parses IPFS log CIDs.
- **Notifications**: Sends notifications (e.g., via Telegram).
- **API Endpoints**:
  - **Exit requests**:
    - `GET /api/v0/event_indexer/exit_requests`
      - query params:
        - `operatorId`: Operator ID to get exit requests.
  - **Operator ID (/)**: Allows for updating the Lido operator ID.
    - `POST api/v0/events_indexer/operatorId`: Updates the operator ID.
      - query params:
        - `operatorId`: Operator ID to update.
    - `DELETE api/v0/events_indexer/operatorId`: Deletes the operator ID.
      - query params:
        - `operatorId`: Operator ID to delete.
  - **Telegram**: Configures Telegram for notifications:
    - `POST /api/v0/events_indexer/telegramConfig`: Configures Telegram.
      - body (JSON):
        - `chatId`: Telegram chat ID.
        - `token`: Telegram bot token.
    - `GET /api/v0/events_indexer/telegramConfig`: Retrieves Telegram configuration.
  - **Validator Performance**: Retrieves validator performance data by operator ID.
    - `GET /api/v0/event_indexer/operator_performance`: Retrieves the operator performance.
      - query params:
        - `operatorId`: Operator ID to get performance data.
- Proxy API: a proxy API that redirect requests to the [Lico API keys](https://github.com/lidofinance/lido-keys-api). Its main functionality is to avoid the cors issues when the Lido CSM UI tries to fetch the API.

## Environment Variables

To configure the app, set the following environment variables:

| Variable         | Description                                                                                          |
|------------------|------------------------------------------------------------------------------------------------------|
| `NETWORK`        | Ethereum network (e.g., `mainnet`, `holesky`). Default holesky                                       |
| `API_PORT`       | Port on which the API will be exposed. Default 8080                                                  |
| `PROXY_API_PORT` | Proxy port on which the Proxy API will be exposed. Default 8081                                                 |
| `BEACONCHAIN_URL`| URL of the Ethereum beacon chain client. Default http://beacon-chain.<network>,dncore.dappnode:3500  |
| `WS_URL`         | URL of the Ethereum WebSocket client. Default ws://execution.<network>.dncore.dappnode:8546          |
| `RPC_URL`        | URL of the Ethereum RPC client. Default http://execution.<network>.dncore.dappnode:8545              |
| `IPFS_URL`       | URL of the IPFS gateway used to fetch logs. Default http://ipfs.dappnode:5001                        |
| `LOG_LEVEL`      | Logging level (e.g., `DEBUG`, `INFO`, `WARN`, `ERROR`). Default INFO                                 |
| `CORS`           | CORS origins domains allowed to fetch the api. Default to Lido CSM ui                                |

Example `.env` file:

```plaintext
# .env
API_PORT=8080
BEACONCHAIN_URL=https://beaconchain.example.com
IPFS_GATEWAY_URL=https://ipfs.example.com
LOG_LEVEL=INFO
```

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/dappnode/lido-events.git
   cd lido-events
    ```

2. **Set Up Environment Variables**: Create a .env file or set environment variables directly.

3. **Install Dependencies**:

    ```bash
    go mod tidy
    ```

4. **Build the Project**:

    ```bash
    go build -o bin/lido-events ./cmd/main.go
    ```

## Running the Application

- Start the App:

    ```bash
    go run ./cmd/main.go
    ```

- Docker Option:

    ```bash
    docker compose up
    ```

## Development

- Project Structure:
  - `/internal`: Core business logic
    - `/adapters`: Interfaces with external dependencies
    - `/application`: Application services, ports, and domain models
    - `/config`: Configuration management
    - `/logger`: Logging setup
  - `/cmd`: Main application entry point

- Adding New Features: Follow hexagonal architecture principles.
- Logging: logging is always done in the service layer, not in the adapter layer.

## Test

- Integration

    ```bash
    WS_URL="" go test -v -tags=integration ./...  
    ```

- Integration with race conditions

    ```bash
    WS_URL="" go test --race -v -tags=integration ./...  
    ```

- Unit

    ```bash
    go test -v  -tags=unit ./...
    ```
