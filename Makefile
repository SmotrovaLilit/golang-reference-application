GOLANGCI_LINT_BIN=golangci-lint

$(GOLANGCI_LINT_BIN):
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: $(GOLANGCI_LINT_BIN)
	$(GOLANGCI_LINT_BIN) run ./...

tests:
	go test -v ./...

generate:
	go generate ./cmd/server/wire.go

build: generate
	go build ./cmd/server

clean:
	go clean -cache