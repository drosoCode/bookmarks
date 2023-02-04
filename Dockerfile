FROM alpine:latest

RUN apk add --update ca-certificates
COPY bookmarks /bookmarks

ENTRYPOINT [ "/bookmarks" ]