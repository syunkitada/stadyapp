gen:
	mkdir -p openapi_schemas
	curl -Lo openapi_schemas/iam.json localhost:8001/openapi.json
	curl -Lo openapi_schemas/compute.json localhost:8002/openapi.json
	npx @hey-api/openapi-ts -i openapi_schemas/iam.json -o src/clients/iam -c @hey-api/client-axios
	npx @hey-api/openapi-ts -i openapi_schemas/compute.json -o src/clients/compute -c @hey-api/client-axios
