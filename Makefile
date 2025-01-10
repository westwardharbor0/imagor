.PHONY: build

prepare: # Tidy and install the cli command.
	go mod tidy
	go install ./cmd/cli

build: # Build a binary file to /build/bin/.
	go build -o ./build/bin/imagor ./cmd/cli
