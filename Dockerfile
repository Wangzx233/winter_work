FROM golang as build

ENV GOPROXY=https://goproxy.io

ADD . /usr/local/go/src/bibi

WORKDIR /usr/local/go/src/bibi

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server


FROM alpine:3.7

ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV MysqlDSN=""
ENV GIN_MODE="release"
ENV PORT=8080

RUN  apk add ca-certificates





WORKDIR /www

COPY --from=build /usr/local/go/src/bibi/api_server /usr/bin/api_server


RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]