FROM golang:latest 

WORKDIR /app

COPY ./go.mod ./go.sum /app/

RUN go mod download

COPY ./ /app/

RUN go build -o user .

EXPOSE 9002

CMD ["./user"]
