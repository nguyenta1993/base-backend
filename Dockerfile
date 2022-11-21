FROM golang:1.19-alpine AS build

# git is required to fetch go dependencies
RUN apk add --no-cache ca-certificates git

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN GOPATH=/Users/$USER/go
RUN export PATH=$GOPATH/bin:$PATH

COPY . .

RUN swag init --parseDependency --parseInternal
RUN go build -o base-backend .

FROM alpine:3.16

WORKDIR /app
COPY --from=build /app/docs ./docs
COPY --from=build /app/base-backend .
COPY --from=build /app/config ./config
COPY --from=build /app/internal/resources ./internal/resources

EXPOSE 5000 5001
CMD ./base-backend start
