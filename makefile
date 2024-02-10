IMAGE_NAME = "a-go-rinha"
DOCKER_USER = "nicolasmmb"


build-and-submit:
	@docker buildx build --platform linux/amd64 -t $(DOCKER_USER)/$(IMAGE_NAME):latest . --push
	@docker buildx build --platform linux/arm64 -t $(DOCKER_USER)/$(IMAGE_NAME):latest-arm . --push

composer-start:
	@docker-compose -f "docker-composer.yaml" up -d --build --force-recreate --remove-orphans

composer-stop:
	@docker-compose -f "docker-composer.yaml" down --remove-orphans

