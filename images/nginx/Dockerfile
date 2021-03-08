FROM nginx:1.18.0-alpine
WORKDIR /app
COPY ./*.json /app/
COPY ./start.sh /start.sh
COPY nginx.conf /etc/nginx/nginx.conf

ENTRYPOINT [ "/start.sh" ]