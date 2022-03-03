FROM golang:rc-alpine as base

RUN apk update \
    && apk add \
        git

WORKDIR /src

COPY . .

RUN go build

ENTRYPOINT [ "/src/sonarqube-diff" ]
