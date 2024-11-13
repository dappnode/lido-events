# Lido events

Repository for indexing Lido events and notifying subscribers.

## Configuration

- `WS_URL`: Websocket URL to listen for events.
- `IPFS_URL`: IPFS URL to fetch the event data.
- `NETWORK`: Network to listen for events.

## Start 

- Run the service

    ```bash
    WS_URL="" go run main.go
    ```

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
