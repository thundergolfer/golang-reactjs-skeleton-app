FROM golang:1.9 AS builder

RUN apt-get update
RUN apt-get install coreutils

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN mkdir -p /go/src/github.com/***
WORKDIR /go/src/github.com/***

COPY Gopkg.toml Gopkg.lock ./
# copies the Gopkg.toml and Gopkg.lock to WORKDIR

RUN dep ensure -vendor-only
# install the dependencies without checking for go code

# build directories
RUN mkdir /app
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

# Build my app
RUN go build -o backend/ . 
CMD ["/app/main"]
