kubectl apply -f ./k8s-prometheus.yaml
kubectl apply -f ./k8s-app.yaml
kubectl rollout restart deploy/app deploy/prometheus
