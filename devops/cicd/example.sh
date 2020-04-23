branchs=('dev')
for branch in ${branchs[@]}; do
    echo ${branch} #分支名
    #注册
    docker run --rm \
        -v /Users/Shared/gitlab-runner/config:/etc/gitlab-runner \
        gitlab/gitlab-runner:alpine-v12.10.0 register \
        --tag-list "${branch}" \
        --non-interactive \
        --registration-token "__2gbqHicD7eQUZrXTPY" \
        --run-untagged="false" \
        --limit=4 \
        --locked="true" \
        --custom_build_dir-enabled="true" \
        --access-level "not_protected" \
        --name "runner-${branch}" \
        --url "https://gitlab.colipu.com/" \
        --executor "docker" \
        --docker-tlsverify="true" \
        --docker-image "docker:19.03.8" \
        --docker-privileged="true" \
        --docker-pull-policy="if-not-present" \
        --docker-volumes '/certs/client' \
        --docker-volumes '/Users/Shared/gitlab-runner/cache:/cache' \
        --docker-volumes '/Users/Shared/gitlab-runner/config/daemon.json:/etc/docker/daemon.json'
done
