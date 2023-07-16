FROM golang:latest

WORKDIR /

COPY . .

RUN go build

EXPOSE 8080

CMD ["./GO-LOGIN"]
