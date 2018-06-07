#!/usr/bin/env bash

#delete ingress rules
kubectl delete -f ingress-rules-istio.yml

#delete services
kubectl delete -f services.yml

#delete deployments
kubectl delete -f deployments.yml

#delete namespace
kubectl delete -f namespace.yml
