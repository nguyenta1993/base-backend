FROM golang:1.19.2-buster

WORKDIR /app

COPY . .
RUN go build -o /server

EXPOSE 5000 5001 6001

CMD ["/server", "start", "--config=./config/local/config.yaml"]

