run: install
	./bin/server
install:
	go get -d -v ./...
	go build -o ./bin/server 
lint:
	go vet ./...
	staticcheck ./...
format:
	go fmt ./...
test:
	go test ./...
