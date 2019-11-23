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
go-shell
```

### Ejecutar mediante contenedor local

* Construir imagen Docker 

```
docker build -t go-kubernetes .
```

* Tag 

```
docker tag go-kubernetes pablon27/go-hello-world:1.0.0
```

* Run

```
docker run -p 8080:8080 pablon27/go-hello-world:1.0.0
```

* Subir imagen a Docker Hub 

```
docker push pablon27/go-hello-world:1.0.0
```

## Kubernetes 

* Deployment 

```
kubectl apply -f k8s-deployment.yml
```

```
kubectl port-forward go-hello-world-dc8589f96-qdsk7 8080:8080
```
