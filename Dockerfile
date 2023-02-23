# FROM golang:1.15.0-alpine as builder
# RUN mkdir /app
# ADD . /app
# WORKDIR /app
# RUN go clean --modcache
# RUN go mod download
# CMD ["go","run","main.go"]
