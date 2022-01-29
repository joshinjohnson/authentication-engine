BINARY_NAME=authentication-engine

clean:
	go clean ./...
	rm -f $(BINARY_NAME)

test:
	go test ./...

build: clean test
	go build -i -o $(BINARY_NAME) -v cmd/main.go

install: build
	mv $(BINARY_NAME) $(GOPATH)/bin
