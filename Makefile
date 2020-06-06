makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list bootstrap init build test

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

bootstrap:
	@make init
	
init:
	go mod tidy

purge:
	rm -rf $(makefile_dir)/test/output/*.svg

build:
	go build

test:
	@make purge
	go test -v ./...
