FROM golang:latest AS builder

WORKDIR /app

COPY ./go.mod ./go.sum /app/

RUN go mod download

COPY ./ /app/

RUN go build -o admin

FROM alpine:latest

WORKDIR /app 

COPY --from=builder /app/admin /app/

EXPOSE 9001

CMD ["./app/admin"]
