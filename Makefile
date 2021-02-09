run:
	@go run cmd/wisdom.go dispense

test:
	@go test ./...

.PHONY: run test