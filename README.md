# D3C

## Compilation

- `go build agent.go`

## Hide console Windows

- `GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui agent.go`
