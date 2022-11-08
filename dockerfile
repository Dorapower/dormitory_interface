# syntax=docker/dockerfile:1

FROM golang:1.19.1-bullseye

WORKDIR /app
COPY go.mod *go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn
RUN go mod download

COPY *.go ./
COPY auth ./auth
COPY sql ./sql
COPY web ./web
RUN go build -o /main

EXPOSE 8090

CMD [ "/main" ]