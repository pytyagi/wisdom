COVERPROFILE=.cover.out
COVERDIR=.cover
PORT ?=3000

run:
	@go run cmd/wisdom.go dispense

test:
	@go test -coverprofile=$(COVERPROFILE) ./...

cover: test
	@mkdir -p $(COVERDIR)
	@go tool cover -html=$(COVERPROFILE) -o $(COVERDIR)/index.html
	@cd $(COVERDIR) && python -m SimpleHTTPServer $(PORT)

.PHONY: run test

