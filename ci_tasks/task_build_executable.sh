#!/bin/sh

set -e
set -x

export GOPATH=$(pwd)
#/gopath:$(pwd)/gopath/src/github.com/cloudfoundry-community/simple-go-web-app/Godeps/_workspace
#cd gopath/src/github.com/cloudfoundry-community/simple-go-web-app/
ls -la
cd skeleton/$SERVICE_NAME
ls -la
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' $SERVICE_NAME.go
ls -laR ..
cp $SERVICE_NAME ../binaries
# cp version.txt ../binaries
cp Dockerfile ../binaries
if [ ! -f ../binaries/Dockerfile ]; then
  echo "It doesn't appear that given Dockerfile is a file joe"
fi