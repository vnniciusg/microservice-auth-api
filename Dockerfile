FROM golang:1.21.3-alpine3.18

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main cmd/main.go

CMD [ "/app/main" ]

EXPOSE 8080