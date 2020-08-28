package module

import (
	"time"

	itemRepo "github.com/OLTeam-go/sea-store-backend-items/repository"
	uItem "github.com/OLTeam-go/sea-store-backend-items/usecase"
)

type itemUsecase struct {
	repo           itemRepo.Repository
	timeoutContext time.Duration
}

// New function initialize usecase for services
func New(repo itemRepo.Repository, tc time.Duration) uItem.Usecase {
	return &itemUsecase{
		repo:           repo,
		timeoutContext: tc,
	}
}
