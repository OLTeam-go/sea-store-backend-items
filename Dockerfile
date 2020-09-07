FROM golang:1.15.0
RUN apt update && apt upgrade -y
EXPOSE 9090
WORKDIR /app
COPY . .
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init
RUN make engine
CMD ["./server"]
