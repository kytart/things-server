build:
	go build server.go

run:
	go run server.go

test:
	go test ./...

clean:
	rm -rf server
