#!/bin/bash
set -e

cd $GOPATH/src/github.com/KyleBanks/go-kit

go get -u github.com/golang/lint/golint

echo "-------------  TEST  -------------"
go test -cover $@ $(go list ./... | grep -v vendor/)

echo "-------------  VET  -------------"
go vet $(go list ./... | grep -v vendor/)

echo "-------------  LINT  -------------"
golint $(go list ./... | grep -v vendor/)

echo "-------------  FMT  -------------"
go fmt $(go list ./... | grep -v vendor/)
