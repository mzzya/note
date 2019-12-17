kubectl exec -it deploy/busybox nslookup redis-svc
kubectl exec -it deploy/busybox nslookup Web_Listener.it.colipu.com
kubectl exec -it deploy/busybox telnet redis-svc 6379
kubectl exec -it deploy/busybox telnet Web_Listener.it.colipu.com 1433

kubectl exec -it deploy/busybox ping Web_Listener.it.colipu.com


