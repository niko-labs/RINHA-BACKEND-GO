build-server:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server -ldflags="-s -w" ./main.go 
	
build-server-mac:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/server -ldflags="-s -w" ./main.go
