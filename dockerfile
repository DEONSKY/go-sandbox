# syntax=docker/dockerfile:1

FROM golang:1.17.5-alpine

RUN apk update

RUN apk add git curl

WORKDIR /root/src

RUN git clone https://github.com/DEONSKY/go-sandbox.git

WORKDIR /root/src/go-sandbox

RUN rm -rf vendor

# Fetch the dependencies
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh


RUN go build -o /go-sandbox-app

EXPOSE 8080

CMD [ "/go-sandbox-app" ]