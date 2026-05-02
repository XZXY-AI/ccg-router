.PHONY: test build lint vuln

test:
	go test ./... -race -count=1

build:
	CGO_ENABLED=0 go build -trimpath -o dist/ccg-router ./cmd/ccg-router

lint:
	go vet ./...

vuln:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
