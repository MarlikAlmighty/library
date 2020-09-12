FROM golang:1.14-alpine3.12 AS builder

WORKDIR /go/src/library

COPY . .

RUN go build -o /go/src/library/app /go/src/library/cmd/library-server/main.go

FROM alpine:3.12

COPY --from=builder /go/src/library/app /

CMD ["/app"]
