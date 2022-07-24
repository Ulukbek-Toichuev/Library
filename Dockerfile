FROM golang:alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o server cmd/server/main.go


EXPOSE 8080

CMD ["/app/server"]

