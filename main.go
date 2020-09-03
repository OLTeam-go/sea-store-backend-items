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
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Service Items API
// @version 1.0
// @description This is an open api for item service

// @contact.name OLTeamgo API Support
// @contact.email yoganandamahaputra@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host sea-store-backend-items.herokuapp.com
// @BasePath /api/v1

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

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
