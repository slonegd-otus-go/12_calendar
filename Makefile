all: gen test build

gen:
	go generate ./...

test:
	go test ./... -cover
	
build:
	go build -o mycalendar main.go