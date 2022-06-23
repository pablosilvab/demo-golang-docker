# Demo Docker & Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/pablosilvab/demo-golang-docker)](https://goreportcard.com/report/github.com/pablosilvab/demo-golang-docker) [![Build Actions Status](https://github.com/pablosilvab/demo-golang-docker/workflows/build/badge.svg)](https://github.com/pablosilvab/demo-golang-docker/actions)


This is a simple API for testing different tools.
### Run project locally

1. Clone project

```
go get github.com/pablosilvab/golang-docker
```

2. Go to folder

```
cd ~/go/src/github.com/pablosilvab/golang-docker
```

3. Download dependencies

```
go mod download
```

4. Run project locally

```
make go-run
```

5. Compile the code and generate an executable 

```
make go-build
```

6. Run the executable

```
make go-shell
```

### Run with Docker

* Build new image 

```
make docker-build
```

* Generate tag 

```
make docker-tag
```

* Run container

```
make docker-run
```

## Helm 

### Requirements

* Helm
* Tiller


Install package in Kubernetes: 
```
make helm-install
```

Remove release in Kubernetes:
```
make helm-uninstall
```

List all releases:
```
helm ls --all
```

## Issues

### Deploy 

* Cannot connect to the Docker daemon at tcp://localhost:2375. Is the docker daemon running?

Review  ```.gitlab-ci.yml``` file. 

### Helm install

* User "system:serviceaccount:default:default" cannot get at the cluster 

Grant permissions with ```ClusterRoleBinding```.

### How create charts?

```
helm create charts
```
