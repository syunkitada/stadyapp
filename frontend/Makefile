gen:
	mkdir -p openapi_schemas
	npx @hey-api/openapi-ts -i ../backends/iam/internal/iam-api/spec/spec_bundled.yaml -o src/clients/iam -c @hey-api/client-axios
	npx @hey-api/openapi-ts -i ../backends/compute/internal/compute-api/spec/spec_bundled.yaml -o src/clients/compute -c @hey-api/client-axios

run:
	npm run dev
