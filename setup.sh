#!/bin/sh

export GOOGLE_APPLICATION_CREDENTIALS=$PWD/your-service-key.json
export PATH=$PATH:$(go env GOPATH)/bin
export GOPATH=$(go env GOPATH)

go install aip