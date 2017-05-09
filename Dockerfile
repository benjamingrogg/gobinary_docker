FROM scratch

COPY pong /

EXPOSE 8888:8888

ENTRYPOINT ["/pong"]


