FROM golang:1.19.2-buster
WORKDIR /app
COPY . .
RUN go build -o /base_service_server

EXPOSE 5000 5001 6001

CMD ["/base_service_server", "start", "--config=./config/dev/config.yaml"]
