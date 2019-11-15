kubectl config use-context minikube

kubectl apply -f ./cluster_role.yaml

kubectl apply -f ./role_binding.yaml

kubectl get rolebindings -A

kubectl describe rolebinding gitlab-runner-rb