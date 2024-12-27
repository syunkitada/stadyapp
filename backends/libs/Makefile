env:
	@make install
	@make gen
	@make db

install:
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2

test:
	go test -v ./...

lint:
	golangci-lint run ./...
