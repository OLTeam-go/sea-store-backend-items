BINARY=server
engine:
	go build -o ${BINARY} main.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -f Dockerfile -t sea-store-backend-items .

run:
	go run main.go

stop:
	docker-compose stop

test: 
	go test -v -cover -covermode=atomic ./...

.PHONY: clean install engine docker run stop test
