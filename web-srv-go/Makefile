# Variables
NAME=go-web-server
REPO=sundrop/$(NAME)
VERSION_FILE=VERSION
VERSION=$(shell cat $(VERSION_FILE))
CONTAINER_NAME=go-srv
LOCAL_PORT=8888

rebuild:
	docker build --no-cache -t $(NAME):$(VERSION) -f docker/Dockerfile .

build:
	docker build -t $(NAME):$(VERSION) -f docker/Dockerfile .

run: stop rm build
	docker run -d -p $(LOCAL_PORT):80 --name $(CONTAINER_NAME) $(NAME):$(VERSION)

stop:
	docker kill $(CONTAINER_NAME) || true

rm:
	docker rm $(CONTAINER_NAME) || true

tag: bump-version build
	@$(eval VERSION=$(shell cat $(VERSION_FILE)))
	docker tag $(NAME):$(VERSION) $(REPO):$(VERSION)
	docker push $(REPO):$(VERSION)

bump-version:
	@echo $(VERSION) | awk -F. '{printf("%d.%d.%d", $$1, $$2, $$3+1)}' > $(VERSION_FILE)

git-tag:
	git add $(VERSION_FILE)
	git commit -m "Bump version to $(VERSION)"
	git tag -a "v$(VERSION)" -m "Version $(VERSION)"
	git push

push:
	docker push $(REPO):$(VERSION)
	git push

print-version:
	@echo $(VERSION)