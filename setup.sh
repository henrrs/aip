#!/bin/sh

export GOOGLE_APPLICATION_CREDENTIALS=/home/development/shared/Projects/aip/outros/service-account-key.json
export PATH=$PATH:$(go env GOPATH)/bin
export GOPATH=$(go env GOPATH)

go install aip