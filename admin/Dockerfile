FROM golang:latest 

WORKDIR /app

COPY ./go.mod ./go.sum /app/

RUN go mod download

COPY ./ /app/

RUN go build -o admin .

EXPOSE 9001

CMD ["./admin"]
