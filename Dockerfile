FROM golang:alpine as build
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o main main.go

EXPOSE 8080
ENTRYPOINT ["./main"]