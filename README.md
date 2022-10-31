# Base Service

## Overview
A sample service to demonstrate the structures and use cases of some common 3rd party tools.

## Prerequisites
- [Go] (https://golang.org/dl/) 1.19 or later
- [Docker] (https://docs.docker.com/get-docker/) 20.10 or later

## Getting Started

### Running the service
1. Clone the repository
2. Run `make run` to start the service
3. Run `make test` to run the tests.
4. Run `make migrate` to create migration.

We have 3 config folder for different environment. You can change the config file in `make run` to run the service in different environment.
Default configuration is set to run on port 5001 for http, 5002 for grpc. you can change it in the `config.yaml` file respectively.

### Docker compose
```bash
$ make compose
```
This will use docker-compose.local.yaml to build up environment

### Build
```bash
$ make build
```
### Docker build
```bash
$ make docker-build
```

## Tools
### Open Telemetry - Jaeger
Distributed tracing system
(https://www.jaegertracing.io/)
### Prometheus
Prometheus exporter for hardware and OS metrics exposed by *NIX kernels, written in Go with pluggable metric collectors.
(https://github.com/prometheus)
### Grafana
Grafana is a multi-platform open source analytics and interactive visualization web application. It provides charts, graphs, and alerts for the web when connected to supported data sources
(https://grafana.com/)
## API
### Swagger
```bash
$ make swagger
```
### Swagger UI
```link
http://localhost:5001/swagger/index.html
```
Default username/password is `admin`/`admin`


### Project Structure
```
├── Makefile    # some cmd utilitis
├── README.md   # This file
├── config      # Config folder
│   ├── config.go   Config instance
│   ├── dev
│   │   ├── config.yaml     #Config value for dev env
│   │   └── prometheus.yml  
│   ├── local
│   │   ├── config.yaml     #.... for local
│   │   └── prometheus.yml
│   └── prod                      #... for prod
│       ├── config.yaml
│       └── prometheus.yml
├── database
│   └── database.go  #Helper to connect to database using sqlx
├── deployment  #Docker file for each environment
│   ├── dev
│   │   └── Dockerfile
│   ├── local
│   │   └── Dockerfile
│   └── prod
│       └── Dockerfile
├── docker-compose.dev.yaml     
├── docker-compose.local.yaml   # Build up infrastructure for local development
├── docs                        # Swagger folder 
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── install.sh
├── internal # include all business logic 
│   ├── api
│   │   ├── container.go # A simple container for DI
│   │   ├── grpc  #GRPC Server 
│   │   │   ├── grpc_server.go #Build up the server 
│   │   │   ├── proto #Example protobuf, this should be move to a common repository
│   │   │   │   ├── user.proto
│   │   │   │   └── user_messages.proto
│   │   │   ├── proto_gen
│   │   │   │   ├── user.pb.go
│   │   │   │   ├── user_grpc.pb.go
│   │   │   │   └── user_messages.pb.go
│   │   │   └── user_grpc.go #Implement GRPC server
│   │   ├── http #HTTP Server
│   │   │   ├── http_server.go #Build up
│   │   │   └── v1
│   │   │       ├── routes.go #Map route to handler 
│   │   │       └── user_handler.go #Map http handler to business logic handler, validation happen here
│   │   └── kafka  #Kafka consumer, implement worker to handle kafka message
│   │       ├── consumer.go
│   │       ├── proto #Sample proto as always
│   │       │   └── kafka_user_messages.proto
│   │       ├── proto_gen
│   │       │   └── kafka_user_messages.pb.go
│   │       └── user_consumer.go #Concrete consumer to handle the message
│   ├── application  #Core business logic 
│   │   └── user #Related domain 
│   │       ├── commands  #Commands mean this will alter the state, e.g UPDATE, DELETE, INSERT stuff
│   │       │   ├── create_user
│   │       │   │   ├── create_user_command.go  #Command object, contain information
│   │       │   │   └── create_user_handler.go  #Handle the command object above
│   │       │   └── update_user
│   │       │       ├── update_user_command.go
│   │       │       └── update_user_handler.go
│   │       ├── queries # Queries just fetch the data 
│   │       │   └── get_user
│   │       │       ├── get_user_handler.go
│   │       │       ├── get_user_handler_test.go
│   │       │       ├── get_user_query.go
│   │       │       └── user.go   #You can transform the data here 
│   │       └── user_service.go  #A container for user's commands and queries, optional
│   ├── domain  #Business Domain, inclue entities and interface contract
│   │   ├── entities
│   │   │   ├── base.go   #Simple base entity, common property for concrete entity
│   │   │   └── user.go
│   │   └── interfaces  #Interfaces for entitines 
│   │       └── user
│   │           ├── cache_repository.go #Contract with cache like redis, memcache, ...
│   │           ├── command_repository.go #Contract with command repo I.E Write repository
│   │           ├── mocks
│   │           │   └── mock.go
│   │           └── query_repository.go #Contract with query repo I.E Read Repository
│   ├── infrastructure
│   │   └── persistent Concrete implement of repository.
│   │       └── user
│   │           ├── command_repository.go
│   │           ├── query_repository.go
│   │           └── redis_repository.go
│   ├── metrics
│   │   ├── grpc
│   │   │   └── metrics.go
│   │   └── http
│   │       └── metrics.go
│   ├── resources #i18n, translation stuff if needed
│   │   ├── en.json
│   │   ├── fr.json
│   │   └── vi.json
│   ├── service #Another container for all business logic, optional if not needed
│   │   └── service.go
│   ├── validation
│   │   └── custom_validation.go #Custom validation rules 
│   ├── wire.go  #Dependency Injection tools
│   └── wire_gen.go
├── main.go # Simple entry point
├── migrations #Database Migration file 
│   ├── 20210912132051_create-user-table.down.sql
│   ├── 20210912132051_create-user-table.up.sql
│   ├── 20211006210144_create-product-table.down.sql
│   └── 20211006210144_create-product-table.up.sql
├── pkg
│   ├── constants #App constants.
│   │   └── constants.go
│   ├── healthcheck
│   │   └── healthcheck.go
│   ├── metrics
│   │   └── metrics.go
│   └── string_utils
│       ├── string_utils.go
│       └── string_utils_test.go
└── startup # Config and bootstrap service here.
    └── startup.go
```
