package usecase

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-items/models"
)

// Usecase are the business logic used for this service
type Usecase interface {
	StoreItem(c context.Context, it *models.Item) (*models.Item, error)
	DeleteItem(c context.Context, id string) error
	UpdateItem(c context.Context, id string, it *models.Item) (*models.Item, error)
	GetByID(c context.Context, id string) (*models.Item, error)
	GetByMerchantID(c context.Context, merchantID string, page int) (*[]models.Item, error)
	Fetch(c context.Context, page int) (*[]models.Item, error)
}
