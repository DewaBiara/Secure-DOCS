FROM golang:1.20.2-alpine3.17 as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o /bin/main -v ./cmd/main

FROM alpine:latest

COPY --from=builder /bin/main .

CMD ["./main"]
