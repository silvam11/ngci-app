#!/usr/bin/env bash

#create/apply namespace
kubectl apply -f ngci-namespace.yml

#create/apply deployments
kubectl apply -f <(istioctl kube-inject -f ngci-deployments.yml)

#Create an endpoint to AWX
kubectl create -f ngci-endpoints.yml

#create/apply services
kubectl apply -f ngci-services.yml

#create/apply gateway rules
istioctl create -f ngci-gateway.yml

#create/apply virtual service rules
istioctl create -f ngci-virtual-service.yml

#Create a secret
#kubectl create -n istio-system secret tls istio-ingressgateway-certs --key /etc/istio/ingressgateway-certs/tls.key --cert /etc/istio/ingressgateway-certs/tls.crt

#Create a service entry
istioctl create -f ngci-service-entry.yml