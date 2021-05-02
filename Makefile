test:
	go test ./...
install:
	go get -v ./...
run:
	go run main.go

build:
	go build main.go

docker-up:
	docker-compose -f resources/docker-compose.yml up -d