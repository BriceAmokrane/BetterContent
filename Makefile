build:
	go build -o bin/ cmd/main.go
	./bin/main.exe

run:
	go run cmd/main.go

test:
	go test -v ./...