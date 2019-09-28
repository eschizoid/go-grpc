FROM golang

COPY . /go/src/github.com/eschizoid/go-grpc

RUN go install github.com/eschizoid/go-grpc/server

ENTRYPOINT ["/go/bin/server"]
EXPOSE 5300
