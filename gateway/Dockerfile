FROM golang:latest AS builder

WORKDIR /app

COPY ./go.mod ./go.sum /app/

RUN go mod download

COPY ./ /app/

RUN go build -o gateway .

FROM alpine:latest

WORKDIR /app 

COPY --from=builder /app/gateway /app/

EXPOSE 9000

CMD ["./gateway"]
