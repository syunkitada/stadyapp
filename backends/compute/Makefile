env:
	@make install
	@make gen
	@make db

install:
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2

gen:
	mkdir -p internal/compute-api/spec/oapi
	cd internal/compute-api/spec && npx @redocly/cli lint spec.yaml
	cd internal/compute-api/spec && npx @redocly/cli bundle spec.yaml -o spec_bundled.yaml
	oapi-codegen -generate "types" -package oapi internal/compute-api/spec/spec_bundled.yaml > internal/compute-api/spec/oapi/types.gen.go
	oapi-codegen -generate "client" -package oapi internal/compute-api/spec/spec_bundled.yaml > internal/compute-api/spec/oapi/client.gen.go
	oapi-codegen -generate "server" -package oapi internal/compute-api/spec/spec_bundled.yaml > internal/compute-api/spec/oapi/server.gen.go
	oapi-codegen -generate "spec" -package oapi internal/compute-api/spec/spec_bundled.yaml > internal/compute-api/spec/oapi/spec.gen.go
	go generate ./internal/...

db:
	go run cmd/compute-admin-ctl/main.go

test:
	go test -v ./internal/...

lint:
	golangci-lint run ./...

mysql:
	sudo docker exec -it stadyapp-mysql-1 mysql -uadmin -padminpass

mod:
	go mod tidy
