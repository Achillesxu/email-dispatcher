FROM golang:buster as builder

WORKDIR /go/src/app

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

CMD go run ./cmd/*.go





