# 启动gitlab
sudo docker run --detach \
  --hostname gitlab.example.com \
  --publish 443:443 --publish 80:80 --publish 22:22 \
  --name gitlab \
  --restart always \
  --volume ~/gitlab/config:/etc/gitlab \
  --volume ~/gitlab/logs:/var/log/gitlab \
  --volume ~/gitlab/data:/var/opt/gitlab \
  gitlab/gitlab-ce:12.4.3-ce.0


  GROUP_NAME="smb" #组名

#部署runner
    docker run -d --name gitlab-runner-${GROUP_NAME} --privileged --restart always \
        -v /var/run/docker.sock:/var/run/docker.sock \
        -v /Users/Shared/gitlab-runner-${GROUP_NAME}/config:/etc/gitlab-runner \
        gitlab/gitlab-runner:alpine-v12.4.1

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
    --name "smb-runner" \
    --url "https://gitlab.example.com/" \
    --executor "docker" \
    --docker-tlsverify="true" \
    --docker-image "docker:19.03.4" \
    --docker-privileged="true" \
    --docker-volumes '/certs/client' \
    --docker-volumes '/cache'
done