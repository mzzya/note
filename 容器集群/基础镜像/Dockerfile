FROM alpine:3.10.1

ENV TZ=Asia/Shanghai

RUN cat /etc/apk/repositories \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories\
    && cat /etc/apk/repositories \
    && apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates \
    && apk add --no-cache bash \
    bash-doc \
    bash-completion \
    --no-cache tzdata \
    && rm -rf /var/cache/apk/* \
    && /bin/bash \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone
