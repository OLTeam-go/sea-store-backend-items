# How to run
```
# make sure golang dep already installed
dep ensure
go run db/postgresql/migrations/*.go init
go run db/postgresql/migrations/*.go
go run main.go
```