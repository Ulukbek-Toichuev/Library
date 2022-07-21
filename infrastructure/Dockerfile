FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./


RUN go mod download

COPY . ./

RUN go build -o serv main.go

EXPOSE 8080

CMD ["/app/serv"]

FROM nginx:latest
COPY ./infrastructure/default.conf /etc/nginx/conf.d/default.conf

EXPOSE 80