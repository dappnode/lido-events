# Lido events

Repository for indexing Lido events and notifying subscribers.

## Test

- Integration

    ```bash
    WS_URL="" go test -v -tags=integration ./...  
    # race conditions
    go test -race -v

    ```


- Unit