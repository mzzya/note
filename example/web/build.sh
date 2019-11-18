CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./app ./main.go
docker build --no-cache -t hellojqk/app:0.0.1 .

rm ./app

docker save -o 123.zip hellojqk/app:0.0.1 prom/prometheus:v2.14.0

docker rmi $(docker image ls --filter="dangling=true" --format="{{.ID}}")

eval $(minikube docker-env)

docker load -i 123.zip

rm ./123.zip

# minikube启动时区不正确问题
minikube ssh -- date

minikube ssh -- sudo date -u $(date -u +%m%d%H%M%Y.%S)

minikube ssh -- date

kubectl apply -f ./prometheus.yaml
kubectl apply -f ./k8s-app.yaml
kubectl rollout restart deploy/app deploy/prometheus

docker rmi $(docker image ls --filter="dangling=true" --format="{{.ID}}")

docker images
