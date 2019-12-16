
FROM golang:latest as builder
LABEL maintainer="Pablo Silva <pablonicolassilvabravo@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /app/main .
EXPOSE 8080
ENTRYPOINT ["./main"] 