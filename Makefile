APP_NAME = demo-golang-docker
APP_VERSION = 0.0.1.SNAPSHOT
USER_HUB = pablon27
PORT = 8080

helm-uninstall:
	helm uninstall demo-golang-docker

helm-upgrade:
	helm upgrade demo-golang-docker  ./charts

helm-install:
	helm install --set name=demo-golang-docker demo-golang-docker ./charts
 
deploy-local:
	kubectl run hello-world-golang --image=${USER_HUB}/${APP_NAME}:${APP_VERSION} --restart=Never --port=${PORT}

docker-push:
	docker build -t ${USER_HUB}/${APP_NAME}:${APP_VERSION} .
	docker push ${USER_HUB}/${APP_NAME}:${APP_VERSION}

docker-build: 
	docker build -t ${USER_HUB}/${APP_NAME}:${APP_VERSION} .

docker-run:
	docker run -p ${PORT}:8080 ${USER_HUB}/${APP_NAME}:${APP_VERSION}

go-run:
	go run cmd/main.go

go-build:
	go build -o ./build/${APP_NAME} ./cmd/main.go

go-shell:
	./build/${APP_NAME}


