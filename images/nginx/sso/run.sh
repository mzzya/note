#! /bin/bash

docker rm -f web
docker build -t web .

# docker run -itd -p 8081:80 --name web web
docker run -it -p 8081:80 -e APP_ENV=prd -e DOMAIN=localhost -e COOKIE_TOKEN_NAME=test_clp_tk --name web web
