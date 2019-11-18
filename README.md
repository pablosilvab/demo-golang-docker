# Demo Docker & Golang

## Docker 

* Construir imagen Docker 

```
docker build -t go-kubernetes .
```

* Tag 

```
docker tag go-kubernetes pablon27/go-hello-world:1.0.0
```

* Subir imagen 

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