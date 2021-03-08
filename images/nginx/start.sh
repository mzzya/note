#! /bin/sh
# cat index.json

cp $APP_ENV.json index.json

# cat $APP_ENV.json
cat index.json
/docker-entrypoint.sh nginx -g 'daemon off;'
