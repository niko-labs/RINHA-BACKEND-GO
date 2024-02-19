build-server:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o bin/server -ldflags="-s -w" ./cmd/main.go 

