FROM nginx:1.19.8-alpine
WORKDIR /app
COPY ./*.html /app/
COPY ./start.sh /start.sh
COPY nginx.conf /etc/nginx/nginx.conf

ENTRYPOINT [ "/start.sh" ]