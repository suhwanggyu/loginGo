FROM golang:1.16.2-buster
WORKDIR /usr/src/app
COPY . .
RUN go build main.go
CMD ["./main"]
