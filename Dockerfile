FROM golang:1.18.3-alpine

LABEL maintainer="Lucas Noga <luc4snoga@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o corpos-christie .

ENTRYPOINT ["./corpos-christie"]