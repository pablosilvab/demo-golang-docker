APP_NAME = golang-docker
APP_VERSION = 0.0.1
USER_HUB = pablon27

docker-push:
	docker push ${USER_HUB}/${APP_NAME}:${APP_VERSION}

docker-build: 
	docker build -t ${USER_HUB}/${APP_NAME}:${APP_VERSION} .

docker-run:
	docker run -p 8080:8080 ${USER_HUB}/${APP_NAME}:${APP_VERSION}

go-run:
	go run cmd/main.go

go-build:
	go build -o ./build/${APP_NAME}

go-shell:
	./build/${APP_NAME}
