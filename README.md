# Demo Docker & Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/pablosilvab/demo-golang-docker)](https://goreportcard.com/report/github.com/pablosilvab/demo-golang-docker)

El objetivo de este proyecto es realizar pruebas con distintas herramientas de CI/CD para una aplicación desarrollada en Go.

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

## Helm 

### Requisitos

* Tener Helm instaldo en tu máquina.
* Tener Tiller instalado en tu cluster.
* El usuario debe tener permisos para poder desplegar.


Instalar paquete en Kubernetes: 
```
helm install --name demo-golang-docker ./charts/
helm install --replace --name demo-golang-docker ./charts 
```

Eliminar release de Kubernetes:
```
helm delete demo-golang-docker
```

Reinstalar release:
```
helm delete demo-golang-docker && helm install --replace --name demo-golang-docker
```

Listar releases:
```
helm ls --all
```

## Issues

### Deploy 

* Cannot connect to the Docker daemon at tcp://localhost:2375. Is the docker daemon running?

Revisar archivo ```.gitlab-ci.yml```. Las versiones son importantes, tanto la imagen Docker como las variables definidas.

### Helm install

* User "system:serviceaccount:default:default" cannot get at the cluster 

Conceder permisos con un ```ClusterRoleBinding```.

### Creación de charts

```
helm create charts
```