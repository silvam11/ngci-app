#!/usr/bin/env bash

#Create a service entry
istioctl delete -f ngci-service-entry.yml

##delete virtual service rules
istioctl delete -f ngci-virtual-service.yml
#
##delete gateway rules
istioctl delete -f ngci-gateway.yml

#delete services
kubectl delete -f ngci-services.yml

#delete deployments
kubectl delete -f ngci-deployments.yml

#delete namespace
kubectl delete -f ngci-namespace.yml
