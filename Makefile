BINARY=items-backend
VERSION=0.1
engine:
	go build -o ${BINARY} main.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t komangyogananda/sea-store-backend-items:${VERSION} .

run:
	docker-compose up --build -d

stop:
	docker-compose down