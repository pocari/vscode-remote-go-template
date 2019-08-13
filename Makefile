BIN_NAME=bin/server

.PHONY:
start:
	realize start --server

build:
	go build -o $(BIN_NAME) src/server.go
