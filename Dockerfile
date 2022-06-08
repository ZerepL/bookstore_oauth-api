FROM golang:1.17-alpine as builder

WORKDIR /app

COPY . ./

RUN go mod download

WORKDIR /app/src

RUN go build -o /bookstore_oauth-api

FROM alpine:3.13.6

COPY --from=builder /bookstore_oauth-api /bookstore_oauth-api

EXPOSE 8080

CMD ["/bookstore_oauth-api"]