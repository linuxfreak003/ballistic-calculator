# syntax=docker/dockerfile:1

FROM golang:1.21 AS builder
RUN apt install gcc g++

ENV APP /app/ballistic-calculator

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o "$APP"

CMD $APP
EXPOSE 8080
