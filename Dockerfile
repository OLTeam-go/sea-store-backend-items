FROM golang:1.15.0
RUN apt update && apt upgrade -y
EXPOSE 9090
WORKDIR /app
COPY .env .env
COPY . .
RUN make engine
