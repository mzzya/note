# nginx部署问题

## 反向代理域名配置

```conf
location / {
      proxy_set_header Host testapi.example.com;
      proxy_set_header Origin testapi.example.com;
      proxy_ssl_server_name on;
      proxy_pass https://testapi.example.com;
      # root /app;
}
```
