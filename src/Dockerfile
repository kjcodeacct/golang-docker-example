FROM golang:alpine AS builder

WORKDIR $GOPATH/src/github.com/kjcodeacct/golang_docker_example/src
COPY . .

RUN export CGO_ENABLED=0 && go build -o /go/bin/golang_docker_example

ENTRYPOINT ["/go/bin/golang_docker_example"]