SHELL := /bin/bash

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## example1: run the program with the standard log
.PHONY: example1
example1:
	@go run ./std-log/main.go

## example2: run the program with the new slog
.PHONY: example2
example2:
	@go run ./std-slog/main.go

## example3: run the program with the custom slog
.PHONY: example3
example3:
	@go run ./custom-slog/main.go

## example4: run the program with the text handler
.PHONY: example4
example4:
	@go run ./text-handler/main.go

## example5: run the program with the JSON handler
.PHONY: example5
example5:
	@go run ./json-handler/main.go

## example6: run the program with contextual logs
.PHONY: example6
example6:
	@go run ./contextual-slog/main.go

## example7: run the program with unbalanced keys
.PHONY: example7
example7:
	@go run ./unbalanced-keys/main.go

## example8: run the program with strongly typed keys
.PHONY: example8
example8:
	@go run ./strongly-slog/main.go

## example9: run the program with attributes
.PHONY: example9
example9:
	@go run ./slog-attr/main.go

## example10: run the program with grouping attributes
.PHONY: example10
example10:
	@go run ./grouping-attr/main.go

## example11: run the program with child logger
.PHONY: example11
example11:
	@go run ./child-slog/main.go

## example12: run the program with custom options
.PHONY: example12
example12:
	@go run ./custom-options/main.go

## example13: run the program with hidden values
.PHONY: example13
example13:
	@go run ./hidden-values/main.go
