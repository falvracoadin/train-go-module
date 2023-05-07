#build-arg
ARG PROJECT
ARG REGISTRY01

FROM ${REGISTRY01}/${PROJECT}:builder as builder
##FROM golang:1.15-alpine as builder

Maintainer Mamat Nurahmat <mamat@loyalto.id>

#build-arg
ARG BRANCH
ARG PORT
ARG PROJECT

# Utilities
##RUN apk add git gcc g++ tzdata zip ca-certificates
# Add dep for package management
##RUN go get -u -f -v github.com/golang/dep/...

#set workdir
##RUN mkdir -p /go/src/${PROJECT}
#WORKDIR /go/src/${PROJECT}

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

# final stage
FROM alpine:latest
ARG BRANCH
ARG PORT
ARG PROJECT

#Env  
ENV TIMEZONE Asia/Jakarta

#set timezone
RUN apk --no-cache add tzdata && echo "Asia/Jakarta" > /etc/timezone
RUN apk add --update tzdata && \
cp /usr/share/zoneinfo/${TIMEZONE} /etc/localtime && \
echo "${TIMEZONE}" > /etc/timezone && apk del tzdata

#expose
EXPOSE ${PORT}
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/${PROJECT} .

#set env file
COPY .env.${BRANCH} .env

RUN printf "#!/bin/sh\n\nwhile true; do\n\techo \"[INFO] Starting Service at \$(date)\"\n\t(./${PROJECT} >> ./history.log || echo \"[ERROR] Restarting Service at \$(date)\")\ndone" > run.sh
RUN printf "#!/bin/sh\n./run.sh & tail -F ./history.log" > up.sh
RUN chmod +x up.sh run.sh
CMD ["./up.sh"]
