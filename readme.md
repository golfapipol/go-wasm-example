# Try go wasm

## Build wasm

```bash
GOARCH=wasm GOOS=js go build -o main.wasm client/main.go
```

## Run server

```bash
go run main.go
```