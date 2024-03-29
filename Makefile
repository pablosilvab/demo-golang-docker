IMAGE_NAME = demo-golang-docker
IMAGE_VERSION = 0.0.1
IMAGE_REGISTRY = pablon27
PORT = 8080
IMAGE_TAG=$(IMAGE_NAME):$(IMAGE_VERSION)

helm-uninstall:
	helm uninstall demo-golang-docker

helm-upgrade:
	helm upgrade demo-golang-docker  ./charts

helm-install:
	helm install demo-golang-docker ./charts

docker-push:
	docker build -t ${IMAGE_TAG} .
	docker tag $(IMAGE_TAG) $(IMAGE_REGISTRY)/$(IMAGE_TAG)
	docker push ${IMAGE_REGISTRY}/${IMAGE_TAG}

docker-build: 
	docker build -t ${IMAGE_TAG} .

docker-run:
	docker build -t ${IMAGE_TAG} .
	docker run -p ${PORT}:8080 ${IMAGE_TAG}

go-run:
	go run cmd/main.go

go-build:
	go build -o ./build/${IMAGE_NAME} ./cmd/main.go

go-shell:
	./build/${IMAGE_NAME}

go-download:
	export GO111MODULE=on
	go mod download

# Run in a local cluster (I use Minikube)
deploy-local:
	kubectl run ${IMAGE_NAME} --image=${IMAGE_REGISTRY}/${IMAGE_TAG} --restart=Never --port=${PORT}

expose-local:
	kubectl expose pod ${IMAGE_NAME} --type=LoadBalancer 