FROM golang:1.22.0

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main ./cmd/app/main.go

EXPOSE 8080

CMD ["./main"]
