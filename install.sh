#!/bin/bash
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
export PATH="$PATH:$(go env GOPATH)/bin"
go install github.com/google/wire/cmd/wire@latest
go install github.com/golang/mock/mockgen@v1.6.0
