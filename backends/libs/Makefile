env:
	@make install
	@make gen
	@make db

install:
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2

gen:
	mkdir -p examples/echo_server/spec/oapi
	cd examples/echo_server/spec && npx @redocly/cli lint spec.yaml
	cd examples/echo_server/spec && npx @redocly/cli bundle spec.yaml -o spec_bundled.yaml
	oapi-codegen -generate "types" -package oapi examples/echo_server/spec/spec_bundled.yaml > examples/echo_server/spec/oapi/types.gen.go
	oapi-codegen -generate "client" -package oapi examples/echo_server/spec/spec_bundled.yaml > examples/echo_server/spec/oapi/client.gen.go
	oapi-codegen -generate "server" -package oapi examples/echo_server/spec/spec_bundled.yaml > examples/echo_server/spec/oapi/server.gen.go
	oapi-codegen -generate "spec" -package oapi examples/echo_server/spec/spec_bundled.yaml > examples/echo_server/spec/oapi/spec.gen.go

test:
	go test -v ./...

lint:
	golangci-lint run ./...
