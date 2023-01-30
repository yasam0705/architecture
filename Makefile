-include .env
export

.PHONY: run
run:
	go run cmd/app/main.go


.PHONY: proto-gen
proto-gen:
	protoc protos/*.proto -I. --go_out=.
