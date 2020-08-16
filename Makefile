build:
	go build -o bin/main cmd/server/main.go

clean: 
	rm -rf bin/

run:
	go run cmd/server/main.go
 
fmt:
	go fmt ./...