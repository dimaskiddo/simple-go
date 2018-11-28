GO_OUTPUT ?= main
GO_EXPOSE_PORT ?= 3000
DOCKER_IMAGE_NAME ?= dimaskiddo/api-simple
DOCKER_IMAGE_VERSION ?= latest

go-build:
	dep ensure
	CGO_ENABLED=0 GOOS=linux go build -a -o ./build/$(GO_OUTPUT) main.go

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION) .

docker-run:
	docker run -d -p $(GO_EXPOSE_PORT):$(GO_EXPOSE_PORT) --name $(GO_OUTPUT) --rm $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)

docker-stop:
	docker stop $(GO_OUTPUT)

docker-logs:
	docker logs $(GO_OUTPUT)

clean-go:
	rm -f ./build/$(GO_OUTPUT)

clean-vendor:
	rm -rf ./vendor/*/*

clean-docker:
	docker rmi -f $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)

clean-all:
	docker rmi -f $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION) && rm -f ./build/$(GO_OUTPUT)
