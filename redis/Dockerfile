FROM debian:latest
LABEL author="danya@onlyawp.ru"
LABEL for="avito-tech"

RUN apt update \
    && apt install -y redis-server \
    && apt clean

EXPOSE 6379

CMD ["redis-server", "--protected-mode", "no"]