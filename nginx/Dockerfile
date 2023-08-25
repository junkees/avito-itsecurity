FROM debian:latest
LABEL author="danya@onlyawp.ru"
LABEL for="avito-tech"

RUN apt update \
    && apt install -y nginx \
    && apt clean

COPY ../nginx.conf /etc/nginx/nginx.conf

EXPOSE 8089


CMD ["nginx", "-g", "daemon off;"]