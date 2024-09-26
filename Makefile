build:
	@go build -o bin/fpm

run: build
	@./bin/fpm

test:
	@go test -v ./...

