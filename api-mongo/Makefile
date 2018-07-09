GO_OUTPUT ?= api-mongo
GO_EXPOSE_PORT ?= 3000
DOCKER_IMAGE_NAME ?= dimaskiddo/api-mongo
DOCKER_IMAGE_VERSION ?= latest

build:
	CGO_ENABLED=0 GOOS=linux go build -a -o $(GO_OUTPUT) main.go

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION) .

docker-run:
	docker run -d -p $(GO_EXPOSE_PORT):$(GO_EXPOSE_PORT) --name $(GO_OUTPUT) --rm $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)

docker-stop:
	docker stop $(GO_OUTPUT)

docker-logs:
	docker logs $(GO_OUTPUT)

clean-build:
	rm -f $(GO_OUTPUT)

clean-docker:
	docker rmi -f $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)

clean-all:
	docker rmi -f $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION) && rm -f $(GO_OUTPUT)
