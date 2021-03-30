#! /bin/sh

sed -i "s/#COOKIE_TOKEN_NAME#/${COOKIE_TOKEN_NAME}/g" /etc/nginx/nginx.conf
# sed -i "s/#DOMAIN#/${DOMAIN}/g" /etc/nginx/nginx.conf

/docker-entrypoint.sh nginx -g 'daemon off;'
