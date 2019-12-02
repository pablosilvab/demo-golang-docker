# Demo Docker & Golang

## Docker 

### Ejecutar proyecto de forma local

1. Clonar proyecto en directorio.

```
go get github.com/pablosilvab/golang-docker
```

2. Entrar en directorio 

```
cd ~/go/src/github.com/pablosilvab/golang-docker
```

3. Descargar dependencias 

```
go mod download
```

4. Ejecutar proyecto

```
make go-run
```

5. Generar archivo binario

```
make go-build
```

6. Ejecutar archivo binario

```
make go-shell
```

### Ejecutar mediante contenedor local

* Construir imagen Docker 

```
make docker-build
```

* Tag 

```
make docker-tag
```

* Run

```
make docker-run
```


## Kubernetes 

* Deployment 

```
kubectl apply -f golang-docker.yaml
```

* Expose service

```
kubectl expose deployment golang-docker-deployment --type=LoadBalancer --port=8080
```

* Minikube 

```
minikube service golang-docker-deployment
```