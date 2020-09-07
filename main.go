package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	database "github.com/OLTeam-go/sea-store-backend-items/db"
	dItem "github.com/OLTeam-go/sea-store-backend-items/delivery/rest"
	rItem "github.com/OLTeam-go/sea-store-backend-items/repository/postgresql"
	uItem "github.com/OLTeam-go/sea-store-backend-items/usecase/module"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Service Items API
// @version 1.0
// @description Api Documentation for Service Items

// @contact.name OLTeamgo API Support
// @contact.email yoganandamahaputra@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host sea-store-backend-items.herokuapp.com
// @BasePath /api

func migrations(url string) {
	fmt.Println("starting migrations")
	m, err := migrate.New(
		"file://db/migrations",
		url)
	if err != nil {
		log.Println(err.Error())
	}
	if err := m.Up(); err != nil {
		log.Println(err.Error())
	}
	fmt.Println("migrations done")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	}

	dbURL, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		panic("DATABASE_URL did not exists")
	}
	migrations(dbURL)

	pagesize, err := strconv.Atoi(os.Getenv("PAGESIZE"))
	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	port := os.Getenv("PORT")

	db, err := database.GetInstance()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	_, err = db.Exec("SELECT 1")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	repository := rItem.New(db, pagesize)

	timeoutContext := time.Duration(timeout) * time.Second

	usecase := uItem.New(repository, timeoutContext)
	dItem.New(e, usecase)

	log.Fatal(e.Start(fmt.Sprintf(":%s", port)))

}
