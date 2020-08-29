FROM golang:1.15.0
RUN apt update && apt upgrade -y
EXPOSE 9090
ADD . /go/src/github.com/OLTeam-go/sea-store-backend-items
WORKDIR /go/src/github.com/OLTeam-go/sea-store-backend-items
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN rm -rf vendor
RUN dep ensure
RUN go run db/postgresql/migrations/*.go init
RUN go run db/postgresql/migrations/*.go
RUN go build -o server main.go
CMD ["./server"]
