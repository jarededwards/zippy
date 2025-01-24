FROM alpine:3.18.2

LABEL org.opencontainers.image.source=https://github.com/jarededwards/zippy
LABEL org.opencontainers.image.description="simple golang http server"


COPY zippy /usr/bin

EXPOSE 8080

ENTRYPOINT ["/usr/bin/zippy"]
