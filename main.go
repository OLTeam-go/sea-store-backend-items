package main

import (
	"fmt"
	"log"
	"time"

	dItem "github.com/OLTeam-go/sea-store-backend-items/delivery/rest"
	itemMiddleware "github.com/OLTeam-go/sea-store-backend-items/delivery/rest/middleware"
	rItem "github.com/OLTeam-go/sea-store-backend-items/repository/postgresql"
	uItem "github.com/OLTeam-go/sea-store-backend-items/usecase/module"
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", dbHost, dbPort),
		User:     dbUser,
		Password: dbPass,
		Database: dbName,
	})
	defer db.Close()
	_, err := db.Exec("SELECT 1")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	e := echo.New()
	middL := itemMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	e.Use(middleware.Logger())
	repository := rItem.New(db, viper.GetInt("app.pagesize"))

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	usecase := uItem.New(repository, timeoutContext)
	dItem.New(e, usecase)

	log.Fatal(e.Start(viper.GetString("server.address")))

}
