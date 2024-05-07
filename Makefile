build:
	@go build -o bin/somerpg

run: build
	@bin/somerpg