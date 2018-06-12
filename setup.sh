#!/usr/bin/env bash

#create/apply namespace
kubectl apply -f namespace.yml

#create/apply deployments
kubectl apply -f <(istioctl kube-inject -f deployments.yml)

#create/apply services
kubectl apply -f services.yml

#create/apply gateway rules
kubectl apply -f ngci-gateway.yml

#create/apply virtual service rules
kubectl apply -f ngci-virtual-service.yml