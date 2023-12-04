run: build
	./bin/pigoen

build:
	@go build -o bin/pigoen cmd/main.go