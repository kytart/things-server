build:
	go build -o bin/server src/main.go

run:
	go run src/main.go

clean:
	rm -rf bin
