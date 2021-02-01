#! /bin/bash

kind delete cluster

kind create cluster --config ./txconfig.yaml  --retain

kind get kubeconfig