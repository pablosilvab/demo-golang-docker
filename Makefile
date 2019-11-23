APP_NAME = golang-docker
APP_VERSION = 1.0.0

docker-push:
	docker push pablon27/${APP_NAME}:${APP_VERSION}

docker-build: 
	docker build -t ${APP_NAME} .

docker-tag:
	docker tag ${APP_NAME} pablon27/${APP_NAME}:${APP_VERSION}

docker-run:
	docker run -p 8080:8080 pablon27/${APP_NAME}:${APP_VERSION}

go-run:
	go run main.go

go-build:
	go build -o ./build/${APP_NAME}

go-shell:
	./build/${APP_NAME}

deploy:
	kubectl apply -f k8s-deployment.yml

