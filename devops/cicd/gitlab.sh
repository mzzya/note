# 启动gitlab 密码 root 123123123
sudo docker run --detach \
  --hostname gitlab.example.com \
  --publish 443:443 --publish 80:80 --publish 22:22 \
  --name gitlab \
  --restart always \
  --volume ~/gitlab/config:/etc/gitlab \
  --volume ~/gitlab/logs:/var/log/gitlab \
  --volume ~/gitlab/data:/var/opt/gitlab \
  gitlab/gitlab-ce:12.9.4-ce.0

#部署runner
docker run -d --name gitlab-runner --privileged --restart always \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /Users/Shared/gitlab-runner/config:/etc/gitlab-runner \
  -v /Users/Shared/gitlab-runner/cache:/cache \
  gitlab/gitlab-runner:alpine-v12.10.0

GROUP_NAME="example" #组名
#标准构建分支
branchs=('dev' 'test' 'uat' 'prd')
for branch in ${branchs[@]}; do
  echo ${branch} #分支名
  #注册
  docker run --rm \
    -v /Users/Shared/gitlab-runner-${GROUP_NAME}/config:/etc/gitlab-runner \
    gitlab/gitlab-runner:alpine-v12.4.1 register \
    --tag-list "${GROUP_NAME}-${branch}" \
    --non-interactive \
    --registration-token "your gitlab token" \
    --run-untagged="false" \
    --limit=1 \
    --locked="false" \
    --custom_build_dir-enabled="true" \
    --access-level "not_protected" \
    --name "runner-${branch}" \
    --url "https://gitlab.example.com/" \
    --executor "docker" \
    --docker-tlsverify="true" \
    --docker-image "docker:19.03.8" \
    --docker-privileged="true" \
    --docker-pull-policy = "if-not-present" \
    --docker-volumes '/certs/client' \
    --docker-volumes '/Users/Shared/gitlab-runner/cache:/cache' \
    --docker-volumes '/Users/Shared/gitlab-runner/config/daemon.json:/etc/docker/daemon.json'
done

$()$(
  toml
  [runners.machine]
  MachineOptions = [
  "engine-registry-mirror=https://o40mvhma.mirror.aliyuncs.com/"
  ]
)$()
