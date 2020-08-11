# 创建用户私钥 -----BEGIN RSA PRIVATE KEY-----
openssl genrsa -out wy.key 2048

# 使用用户私钥创建【证书请求签名】pem，在-subj中指定用户和组 -----BEGIN CERTIFICATE REQUEST-----
openssl req -new -key wy.key -out wy.pem -subj "/CN=wy/O=group1"

# 生成用户证书
openssl x509 -req -in wy.pem -CA /Users/wangyang/.minikube/ca.crt -CAkey /Users/wangyang/.minikube/ca.key -CAcreateserial -out wy.crt -days 9999

# 配置 kubectl 用户证书信息
kubectl config set-credentials wy --client-certificate=/Users/wangyang/blog/容器集群/k8s-script-study/rbac/wy.crt  --client-key=/Users/wangyang/blog/容器集群/k8s-script-study/rbac/wy.key

# 配置 kubectl 用户上下文
kubectl config set-context minikube-wy --cluster=minikube --user=wy

cat ~/.kube/config

# 从证书内导出公钥
openssl x509 -outform PEM -in ./ca.pem -pubkey -noout -out ca-pub.pem

# 从证书内到处CSR


# 从https 443端口查看证书信息

openssl s_client -connect  *.*.com:443
