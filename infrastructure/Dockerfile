FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./


RUN go mod download

COPY . ./

RUN go build -o serv cmd/server/main.go

EXPOSE 8080

CMD ["/app/serv"]
