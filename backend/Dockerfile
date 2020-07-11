FROM golang

RUN mkdir /backend

ADD . /backend/

WORKDIR /backend

RUN go mod download

EXPOSE 8080

ENTRYPOINT go run ./cmd/api/main.go
