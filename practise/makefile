.DEFAULT_GOAL := build

.PHONY: fmt vet build clean run

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o practise

run: build
	./practise

clean:
	rm -f practise