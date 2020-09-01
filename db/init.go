package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/go-pg/pg"
	"github.com/joho/godotenv"
)

var singleton *pg.DB
var mt sync.Mutex

func connectDatabase() (*pg.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASS")
	dbName := os.Getenv("POSTGRES_DB")
	dbURL, exist := os.LookupEnv("DATABASE_URL")
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	if exist {
		url = dbURL
	}
	fmt.Println(fmt.Sprintf("connceting to postgres = %s", url))

	opt, err := pg.ParseURL(url)
	db := pg.Connect(opt)
	_, err = db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetInstance return the singleton of db connection
func GetInstance() (*pg.DB, error) {
	if singleton == nil {
		mt.Lock()
		defer mt.Unlock()
		if singleton == nil {
			db, err := connectDatabase()
			if err != nil {
				return nil, err
			}
			singleton = db
			return singleton, nil
		}
	}
	return singleton, nil
}
