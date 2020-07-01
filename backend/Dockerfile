FROM golang

COPY . .

WORKDIR .

RUN go mod download

ENTRYPOINT go run main.go
