env:
	@make install
	@make gen
	@make db

install:
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2

gen:
	mkdir -p internal/iam-api/spec/oapi
	oapi-codegen -generate "types" -package oapi internal/iam-api/spec/spec.yaml > internal/iam-api/spec/oapi/types.gen.go
	oapi-codegen -generate "client" -package oapi internal/iam-api/spec/spec.yaml > internal/iam-api/spec/oapi/client.gen.go
	oapi-codegen -generate "server" -package oapi internal/iam-api/spec/spec.yaml > internal/iam-api/spec/oapi/server.gen.go
	oapi-codegen -generate "spec" -package oapi internal/iam-api/spec/spec.yaml > internal/iam-api/spec/oapi/spec.gen.go
	go generate ./internal/...

db:
	go run cmd/iam-admin-ctl/main.go

test:
	go test -v ./internal/...

lint:
	golangci-lint run ./...

mysql:
	sudo docker exec -it stadyapp-mysql-1 mysql -uadmin -padminpass

mod:
	go mod tidy
