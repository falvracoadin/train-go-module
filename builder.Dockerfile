FROM golang:1.16-alpine as builder

Maintainer Mamat Nurahmat <mamat@loyalto.id>

#build-arg
ARG BRANCH
ARG PORT
ARG PROJECT

# Utilities
RUN apk add git gcc g++ tzdata zip ca-certificates
# Add dep for package management
# RUN go get -u -f -v github.com/golang/dep/...

#set workdir
RUN mkdir -p /go/src/${PROJECT}
WORKDIR /go/src/${PROJECT}

# COPY go.mod and go.sum files to the workspace
COPY go.mod . 
COPY go.sum .

## Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
##RUN go mod tidy

COPY . .
RUN go get
COPY .env.${BRANCH} .env

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/${PROJECT} server.go 
