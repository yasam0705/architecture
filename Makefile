-include .env
export

.PHONY: run
run:
	go run cmd/app/main.go