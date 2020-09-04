FROM golang:1.15.0
RUN apt update && apt upgrade -y
EXPOSE 9090
WORKDIR /app
COPY . .
COPY .env.example .env
RUN make engine
CMD ["./server"]
