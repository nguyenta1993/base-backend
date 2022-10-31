FROM golang:1.17.1-buster

WORKDIR /app

COPY . .
RUN go build -o /server
EXPOSE 5000
CMD ["/server", "start", "--config=./config/prod/config.yaml"]

