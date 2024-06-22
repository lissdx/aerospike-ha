# Aerospike

## RRD-like HTTP service

### Purpose

The service's purpose is to store time series data such as   
Types:
* network 
* bandwidth
* temperatures 
* CPU load   

are linked to the concrete timestamp.  
In general, this means we can store one metric  
for each unique combination of type and timestamp.

### API
###### An [openapi.yaml](openapi/openapi.yaml) file is provided to ensure good integrability.

#### PUT /metrics
- Request body:  
  **MetricNode**  application/json  
  example:
  ```
  {
    "type": "network",
    "timestamp": 1000,
    "metric": 777
  }
  ```
  Constraints:  
    - `type`: enum [network, bandwidth, temperature, cpu]
    - `timestamp`: > 0
  
- Response:
  - successfully stored: 200 OK and OkResponse JSON object
  - on error: 400 and badRequest JSON object

#### GET /metrics
- Query parameters:
  - `end` (optional) - timestamp to close the search range (will be excluded from the result)  
    should be greater than 0 if provided
  - `start` (optional) - timestamp to open the search range (will be included in the result)  
    should be greater than 0 if provided
  - `type` (optional) - metric type to search for ( see request `type` ) 
  - `max_fetch` (optional) - constraints the max fetched metrics (default 1000)  
  Constraints: if `start` and `end` params are provided   
  the `end` param should be greater than `start` 
- Response:
  - successfully fetched: 200 OK and array of **MetricNode**s
  - an empty fetch result: 204 No Content
  - on error: 400 and badRequest JSON object

#### Technical Details
##### API
API is defined in [OpenAPI 3](https://swagger.io/specification/v3/) format in the [openapi.yaml](./openapi/openapi.yaml) file.  
DTOs and service interface code is generated using [oapi-codegen](https://github.com/deepmap/oapi-codegen).

##### HTTP server
[Echo](https://echo.labstack.com/) framework is used to manage routes. API generator has a nice integration with this framework.

### Show Me The Money
There are 2 main files:
- HTTP server and handlers: [aerospike_service.go](internal/service/http/rrd_service/aerospike_service.go)  
  includes the HTTP request handlers
- Aerospike DB driver: [aerospike_driver.go](internal/drivers/cache/aerospike_driver.go)  
  includes DB put/get handlers

### Config
The configuration file is provided: [development.env](configs/development.env)

### How To Run
1. The simplest way is to run it in docker-compose the command:
   ```
    docker-compose -f ./docker-compose/docker-compose-all.yml up -d
   ```
   In this case, both the rrd-server and the Aerospike DB will be started.  
   NOTE: There is a small issue with starting AerospikeDB, so we should wait for 20 seconds   
   before running the rrd-server.
2. We can also run the rrd-server and the Aerospike DB in separate containers.
```
 docker-compose -f docker-compose/docker-compose-aerospike.yml up -d
 docker-compose -f docker-compose/docker-compose-asrrd.yml up -d
```
In this case, make sure that the Aerospike DB is running before  
you start the rrd-service.
