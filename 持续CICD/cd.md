# CI

## 安装gitlab-runner

``` docker
docker run -d --name gitlab-runner --restart always \
   -v /Users/Shared/gitlab-runner/config:/etc/gitlab-runner \
   -v /var/run/docker.sock:/var/run/docker.sock \
   gitlab/gitlab-runner:alpine-v12.0.2
```
