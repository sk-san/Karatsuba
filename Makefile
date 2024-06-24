build:
	@go build -o bin/ka
run: build
	@./bin/ka
test:
	@go test ./... -v