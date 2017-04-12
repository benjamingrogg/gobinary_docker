# FROM alpine:3.5
FROM scratch

COPY pong /

EXPOSE 8888

ENTRYPOINT ["/pong"]


