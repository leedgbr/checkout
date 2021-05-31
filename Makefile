build:
	@go build

test:
	@go test ./...

run:
	@go run .

ci: test build

.PHONY: build test ci