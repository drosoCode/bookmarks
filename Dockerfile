FROM debian:11-slim
RUN apt-get update && apt-get install -y ca-certificates
COPY bookmarks /bookmarks
CMD ["/bookmarks"]