# syntax=docker/dockerfile:1

FROM golang:1.19.1-bullseye

WORKDIR /app
COPY go.mod *go.sum ./
RUN go mod download

COPY *.go ./
COPY sql ./sql
RUN go build -o /main

EXPOSE 80

CMD [ "/main" ]