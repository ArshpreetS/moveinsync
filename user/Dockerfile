FROM golang:latest AS builder

WORKDIR /app

COPY ./go.mod ./go.sum /app/

RUN go mod download

COPY ./ /app/

RUN go build -o user

FROM alpine:latest

WORKDIR /app 

COPY --from=builder /app/user /app/

EXPOSE 9002

CMD ["./app/user"]
