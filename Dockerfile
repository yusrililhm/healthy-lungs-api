FROM golang:1.22.2-alpine3.19 AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:3.19.1

WORKDIR /app

COPY --from=builder /app/main .

CMD [ "/app/main" ]
