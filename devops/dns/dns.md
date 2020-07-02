```
.:53 {
        errors
        health {
           lameduck 5s
        }
        ready
        kubernetes cluster.local in-addr.arpa ip6.arpa {
           pods insecure
           fallthrough in-addr.arpa ip6.arpa
           ttl 30
        }
        prometheus :9153
        forward . /etc/resolv.conf {
            except it.colipu.com
        }
        cache 30
        loop
        reload
        loadbalance
    }
    it.colipu.com:53 {
      log
      errors
      cache 30
      forward . 10.101.101.63 10.10.210.89
    }
```

期望 *.it.colipu.com的域名解析使用 10.101.101.63 10.10.210.89 的配置