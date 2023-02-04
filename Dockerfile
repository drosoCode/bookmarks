FROM golang:alpine

COPY bookmarks /bookmarks
RUN go run github.com/playwright-community/playwright-go/cmd/playwright install --with-deps && chmod +x bookmarks

ENTRYPOINT [ "/bookmarks" ]