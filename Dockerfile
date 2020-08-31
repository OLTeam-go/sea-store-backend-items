FROM golang:1.15.0
RUN apt update && apt upgrade -y
EXPOSE 9090
WORKDIR /app
COPY . .
RUN go run db/postgresql/migrations/*.go init
RUN go run db/postgresql/migrations/*.go
RUN go build -o server main.go
CMD ["./server"]
