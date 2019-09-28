FROM golang

ADD . /go/src/github.com/eschizoid/go-grpc/server

RUN go install github.com/eschizoid/go-grpc/server

ENTRYPOINT ["/go/bin/server"]
EXPOSE 5300
