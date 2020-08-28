package postgresql

import (
	rItems "github.com/OLTeam-go/sea-store-backend-items/repository"
	"github.com/go-pg/pg"
)

type postgresqlRepository struct {
	Conn *pg.DB
}

// New function will create object that represent the repository
func New(Conn *pg.DB) rItems.Repository {
	return &postgresqlRepository{
		Conn,
	}
}
