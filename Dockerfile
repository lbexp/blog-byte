FROM golang:1.21.0

WORKDIR /go/src

RUN go install github.com/air-verse/air@latest

COPY . .

CMD air

