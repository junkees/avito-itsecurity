FROM golang:1.21
LABEL author="danya@onlyawp.ru"
LABEL for="avito-tech"

WORKDIR ./app
COPY ./redis-api ./


RUN go mod download
RUN go build -o redis-api main.go

EXPOSE 8081


ENTRYPOINT ["./redis-api"]