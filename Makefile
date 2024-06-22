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


# RUN aerospike DB
.PHONY: as-db-local-up
as-db-local-up:
	docker-compose -f docker-compose/docker-compose-aerospike.yml up -d

# STOP aerospike DB
.PHONY: as-db-local-down
as-db-local-down:
	docker-compose -f docker-compose/docker-compose-aerospike.yml down

# RUN aerospike rrd server
.PHONY: asrrd-local-up
asrrd-local-up:
	docker-compose -f docker-compose/docker-compose-asrrd.yml up -d

# STOP aerospike rrd server
.PHONY: asrrd-local-down
asrrd-local-down:
	docker-compose -f docker-compose/docker-compose-asrrd.yml down

# watch logs of rrd server
.PHONY: asrrd-logs
asrrd-logs:
	docker-compose -f ./docker-compose/docker-compose-all.yml logs -f as-rrd

# RUN aerospike rrd server AND aerospike DB
# single docker compose
.PHONY: local_all_up
local_all_up:
	docker-compose -f ./docker-compose/docker-compose-all.yml up -d

# STOP aerospike rrd server AND aerospike DB
# single docker compose
.PHONY: local_all_down
local_all_down:
	docker-compose -f ./docker-compose/docker-compose-all.yml down

#.PHONY: as-docker-bash
#as-docker-bash:
#	docker exec -ti aerospike /bin/bash