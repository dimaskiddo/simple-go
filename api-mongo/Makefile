GO_OUTPUT ?= main
GO_EXPOSE_PORT ?= 3000
DOCKER_IMAGE_NAME ?= dimaskiddo/api-mongo
DOCKER_IMAGE_VERSION ?= latest

go-build:
	dep ensure
	CGO_ENABLED=0 GOOS=linux go build -a -o ./build/$(GO_OUTPUT) main.go

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION) .

docker-run:
	docker-compose up -d

docker-stop:
	docker-compose down

docker-logs:
	docker-compose logs $(GO_OUTPUT)

clean-build:
	rm -f ./build/$(GO_OUTPUT)

clean-docker:
	docker rmi -f $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)

clean-all:
	docker rmi -f $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION) && rm -f ./build/$(GO_OUTPUT)
