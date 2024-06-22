OPENAPI_FILE_PATH=./openapi/openapi.yaml
GEN_DIR_PATH=internal/pkg/gen/openapi
GEN_PKG_NAME=rrd_server_gen
GEN_FILE_PATH=$(GEN_PKG_NAME)/gen.go

# Install oapi-codegen:
# Ensure you have oapi-codegen installed. If not, you can install it using go get
.PHONY: oapi-codegen-install
oapi-codegen-install:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0


.PHONY: oapi-generate
oapi-generate:
	oapi-codegen -package $(GEN_PKG_NAME) -generate models,server $(OPENAPI_FILE_PATH) > $(GEN_DIR_PATH)/$(GEN_FILE_PATH)


.PHONY: as-local-up
as-local-up:
	docker-compose -f docker-compose/docker-compose-aerospike.yml up -d

.PHONY: as-local-down
as-local-down:
	docker-compose -f docker-compose/docker-compose-aerospike.yml down

.PHONY: as-docker-bash
as-docker-bash:
	docker exec -ti aerospike /bin/bash