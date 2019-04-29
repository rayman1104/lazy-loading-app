FROM golang:latest

WORKDIR /usr/src/app
RUN go get google.golang.org/grpc

COPY . .

EXPOSE 50051
CMD go run main.go
