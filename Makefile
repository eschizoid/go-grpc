GOCMD := go
PROTOCMD := protoc
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
PROTO_OUT := "proto/ingest.pb.go"
PROTO_GW_OUT := "proto/ingestgw.pb.gw.go"
PROTO_SG_OUT := "proto/ingestgw.swagger.json"
SERVER_OUT := "bin/server"
CLIENT_OUT := "bin/client"
SERVER_PKG_BUILD := "github.com/eschizoid/go-grpc/server"
CLIENT_PKG_BUILD := "github.com/eschizoid/go-grpc/client"

.PHONY: all build

all: build-grpc build-grpc-gw build-grpc-sw build

build:
	$(GOBUILD) -i -v -o ${GOPATH}/$(CLIENT_OUT) $(CLIENT_PKG_BUILD)
	$(GOBUILD) -i -v -o ${GOPATH}/$(SERVER_OUT) $(SERVER_PKG_BUILD)

build-grpc:
	$(PROTOCMD) ingest.proto \
	-I./proto \
	--go_out=plugins=grpc:./proto

build-grpc-gw:
	$(PROTOCMD) ingestgw.proto \
	-I./proto \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--plugin=protoc-gen-grpc-gateway=${GOPATH}/bin/protoc-gen-grpc-gateway \
	--grpc-gateway_out=logtostderr=true:./proto

build-grpc-sw:
	$(PROTOCMD) ingestgw.proto \
	-I./proto \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--swagger_out=logtostderr=true:./proto \

test:
	$(GOTEST) -v ./...

clean:
	@rm ${GOPATH}/$(SERVER_OUT) ${GOPATH}/$(CLIENT_OUT) $(PROTO_OUT) $(PROTO_GW_OUT) $(PROTO_SG_OUT)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

dep:
	$(GOCMD) get -v -d ./...
