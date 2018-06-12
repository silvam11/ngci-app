#!/usr/bin/env bash

#echo "Setting K8S docker repo as docker repo ..."
RC=`which minikube`
if [ "${RC" eq 0 ]
then
  echo "Minikube exists then create the images on minikube registry ..."
  eval $(minikube docker-env)
  echo "Done"
fi

echo "Creating docker image for web gateway ..."
docker build -t ngci/web-gateway:0.1 web-gateway

echo "Creating docker image for system model ..."
docker build -t ngci/system-model:0.1 system-model