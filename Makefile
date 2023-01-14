SHELL := bash
#.IGNORE:
#.ONESHELL:
#.SHELLFLAGS := -eu -o pipefail -cli
#.DELETE_ON_ERROR:
#MAKEFLAGS += --warn-undefined-variables
#MAKEFLAGS += --no-builtin-rules

test:
	@echo "Running test endpoint with flags..."
	go run cmd/cli/main.go -dir /Users/james/tmp/smush/ -m file -s log -m file_7 -user logger -pass pickle32 -host 10.63.0.27 -db logtrack -e local.env
	go run cmd/cli/main.go -dir /Users/james/tmp/smush/ -m file -s log -m file_8 -user logger -pass pickle32 -host 10.63.0.27 -db logtrack -e local.env
	go run cmd/cli/main.go -dir /Users/james/tmp/smush/ -s log -user logger -pass pickle32 -host 10.63.0.27 -db logtrack -e local.env

clean:
	@echo "Cleaning test directory..."
	-gunzip /Users/james/tmp/smush/*.gz

build:
	@echo "Building Binary"
	env GOOS=linux GOARCH=amd64 go build -o build/smush-amd64 cmd/cli/main.go  
	env GOOS=darwin GOARCH=arm64 go build -o build/smush-darwin cmd/cli/main.go