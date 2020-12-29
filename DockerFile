FROM golang:latest

RUN mkdir /app
WORKDIR /app

RUN go get github.com/gin-gonic/gin