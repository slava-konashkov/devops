build:
	docker build --no-cache -t my-go-web-server -f docker/Dockerfile .

run: stop rm build
	docker run -d -p 8080:80 --name go-srv my-go-web-server

stop:
	docker kill go-srv || true

rm:
	docker rm go-srv || true