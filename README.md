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