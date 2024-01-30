build:
	@go build -o bin/app ./src

run: build
	@./bin/app

test:
	@go test -v ./tests/...
