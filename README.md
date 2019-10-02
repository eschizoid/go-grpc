#### Protobuf over gRCP
A simple integration between protobuf and gRCP written in Golang.

##### Building
```shell script
$ make clean
$ make all
```

##### Starting gRPC and REST servers
```shell script
$ docker-compose up --build -d
```

##### Testing
```shell script
$ curl -X POST -d '{"message":"ping"}' 'http://go-grpc:8081/v1/ingest'
```

##### Swagger docs
```shell script
$ curl -X GET 'http://go-grpc:8081/swagger/ingestwg.swagger.json'
```

and you should get a response like the following:

```json
{
  "swagger": "2.0",
  "info": {
    "title": "ingestgw.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/ingest": {
      "post": {
        "operationId": "Do",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/protoResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/protoRequest"
            }
          }
        ],
        "tags": [
          "Ingest"
        ]
      }
    }
  },
  "definitions": {
    "protoRequest": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "protoResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    }
  }
}
```