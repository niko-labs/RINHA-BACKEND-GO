build-server:
	@go build -o bin/server -ldflags="-s -w" ./main.go

run-exec:
	@./bin/$(APP_NAME)

build-run: build-server run-exec

