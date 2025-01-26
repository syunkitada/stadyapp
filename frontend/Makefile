gen:
	mkdir -p openapi_schemas
	npx @hey-api/openapi-ts -i ../backends/iam/internal/iam-api/spec/spec_bundled.yaml -o src/clients/iam -c @hey-api/client-axios
	npx @hey-api/openapi-ts -i ../backends/compute/internal/compute-api/spec/spec_bundled.yaml -o src/clients/compute -c @hey-api/client-axios

shadcn:
	npx shadcn@latest add -o sidebar dropdown-menu collapsible avatar breadcrumb table checkbox alert-dialog dialog input button label

run:
	npm run dev
