.PHONY: build run

BUILD_TAG = 1.10.3
IMAGE_NAME = redashbot
IMAGE = $(IMAGE_NAME):$(BUILD_TAG)

build:
	docker build -t $(IMAGE) .

run:
	docker run --rm --name $(IMAGE_NAME) -e SLACK_API_TOKEN $(IMAGE)
