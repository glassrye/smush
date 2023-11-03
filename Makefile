SHELL := bash
#.IGNORE:
#.ONESHELL:
#.SHELLFLAGS := -eu -o pipefail -cli
#.DELETE_ON_ERROR:
#MAKEFLAGS += --warn-undefined-variables
#MAKEFLAGS += --no-builtin-rules

test:
	@echo "Running test endpoint with flags..."
	go run cmd/cli/main.go -dir /Users/james/tmp/smush/ -m file -s log -m file_7 -user logger -pass <pass> -host <host_ip> -db logtrack -e local.env
	go run cmd/cli/main.go -dir /Users/james/tmp/smush/ -m file -s log -m file_8 -user logger -pass <pass> -host <host_ip> -db logtrack -e local.env
	go run cmd/cli/main.go -dir /Users/james/tmp/smush/ -s log -user logger -pass <pass> -host <host_ip> -db logtrack -e local.env

clean:
	@echo "Cleaning test directory..."
	-gunzip /Users/james/tmp/smush/*.gz

binary:
	@echo "Building Binary"
	-env GOOS=linux GOARCH=amd64 go build -o build/smush-amd64 cmd/cli/*.go  
	-env GOOS=darwin GOARCH=arm64 go build -o build/smush-darwin cmd/cli/*.go
