all: gen test build

gen:
	go generate ./...

test:
	go test ./... -cover
	
build:
	go build -o mycalendar main.go

godog:
	docker-compose -f ./docker/docker-compose.yml up  -d ;\
	echo Wait 5s for service availability... ;\
	sleep 5 ;\
	cd tests && godog; \
	test_status_code=$$? ;\
	cd .. ;\
	docker-compose -f ./docker/docker-compose.yml down ;\
	exit $$test_status_code ;\