FROM golang:1.11.1-alpine3.8 AS build-env
WORKDIR /go/src/github.com/SaKu2110/golexa-example
COPY ./ ./
RUN go build -o server main.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/SaKu2110/golexa-example/server /usr/local/bin/server

EXPOSE 8080
CMD ["/usr/local/bin/server"]
