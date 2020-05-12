APP_NAME = demo-golang-docker
APP_VERSION = 0.0.1.SNAPSHOT
USER_HUB = pablon27
PORT = 8080
GIT_DIR=$(shell pwd)

helm-uninstall:
	helm uninstall demo-golang-docker

helm-upgrade:
	helm upgrade demo-golang-docker  ./charts

helm-install:
	helm install --set name=demo-golang-docker demo-golang-docker ./charts
 
docker-push:
	docker build -t ${USER_HUB}/${APP_NAME}:${APP_VERSION} .
	docker push ${USER_HUB}/${APP_NAME}:${APP_VERSION}

docker-build: 
	docker build -t ${APP_NAME}:${APP_VERSION} .

docker-shell:
	docker run -it --rm -v $(GIT_DIR):/app -p ${PORT}:8080  -w /app/ --entrypoint=/bin/sh ${APP_NAME}:${APP_VERSION}

docker-run:
	docker run -p ${PORT}:8080 ${APP_NAME}:${APP_VERSION}

go-run:
	go run cmd/main.go

go-build:
	go build -o ./build/${APP_NAME} ./cmd/main.go

go-shell:
	./build/${APP_NAME}

# To play with Minikube
deploy-local:
	kubectl run hello-world-golang --image=${USER_HUB}/${APP_NAME}:${APP_VERSION} --restart=Never --port=${PORT}

expose-local:
	kubectl expose pod hello-world-golang --type=LoadBalancer 

get-ip-svc:
	minikube service hello-world-golang