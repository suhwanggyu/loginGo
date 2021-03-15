FROM golang:1.62
WORKDIR /usr/src/app
COPY . .
RUN go build main.go
CMD ["main"]