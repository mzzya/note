#! /bin/bash

docker rm -f web
docker build -t web .

# docker run -itd -p 8081:80 --name web web
docker run -it -p 8081:80 -e APP_ENV=test --name web web
