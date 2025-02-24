gen:
	mkdir -p ./internal/myctl/iam
	oapi-codegen -templates templates -generate "types,client,spec" -package iam ../backends/iam/internal/iam-api/spec/spec_bundled.yaml > ./internal/myctl/iam/cli.gen.go
	mkdir -p ./internal/myctl/compute
	oapi-codegen -templates templates -generate "types,client,spec" -package compute ../backends/compute/internal/compute-api/spec/spec_bundled.yaml > ./internal/myctl/compute/cli.gen.go
