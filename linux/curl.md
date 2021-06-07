# curl

## 使用代理

```sh
curl -x 127.0.0.1:8899 your_url
```

如果是HTTPS需要用代理(例如:Whistle)的证书配置添加到环境变量 `export CURL_CA_BUNDLE=~/.curl/rootCA.crt`