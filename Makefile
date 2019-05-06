build:
	@go build ./cmd/sitemap.go

test:
	@go test -short -v ./...

test-race:
	@go test -race -short -v ./...

integration-test:
	@go test ./cmd/...

coverage:
	@go test -cover

lint:
	@golint ./...