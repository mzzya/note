kubectl apply -f ./prometheus-operator.yaml
kubectl create -f ~/kube-prometheus/manifests/setup/
until kubectl get servicemonitors --all-namespaces ; do date; sleep 1; echo ""; done
kubectl create -f  ~/kube-prometheus/manifests/

kubectl delete --ignore-not-found=true -f ~/kube-prometheus/manifests/ -f ~/kube-prometheus/manifests/setup