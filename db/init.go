package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-pg/pg"
)

var singleton *pg.DB
var mt sync.Mutex

func connectDatabase() (*pg.DB, error) {
	log.Println("connecting to database")
	dbURL, exist := os.LookupEnv("DATABASE_URL")
	if !exist {
		panic("DATABASE_URL did not exists")
	}
	fmt.Println(fmt.Sprintf("connceting to postgres = %s", dbURL))

	opt, err := pg.ParseURL(dbURL)
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
