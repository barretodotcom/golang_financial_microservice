FROM golang:latest

WORKDIR /go/src/app

COPY . .

CMD ["go", "run", "main.go"]