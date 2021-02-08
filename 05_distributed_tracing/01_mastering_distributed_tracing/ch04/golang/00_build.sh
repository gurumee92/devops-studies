#!/bin/bash

go build -v -o hello ./service-hello/main.go
go build -v -o bigbrother ./service-bigbrother/main.go
go build -v -o formatter ./service-formatter/main.go

# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./docker/service-hello/main ./service-hello/main.go
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./docker/service-bigbrother/main ./service-bigbrother/main.go
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./docker/service-formatter/main ./service-formatter/main.go
# docker-compose up --build -d