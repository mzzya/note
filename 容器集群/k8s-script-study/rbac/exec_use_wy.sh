#! /bin/bash

kubectl delete -f ./redis.yaml

kubectl config use-context wy-context

kubectl get pods 
kubectl get pods -A

kubectl get nodes
kubectl get nodes -A

kubectl get deployments
kubectl get deployments -A
