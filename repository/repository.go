package repository

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-items/models"
	"github.com/google/uuid"
)

// Repository represent the item repository contract
type Repository interface {
	GetByID(ctx context.Context, id string) (*models.Item, error)
	GetByMerchantID(cxt context.Context, merchantID string, page int) (*[]models.Item, error)
	StoreItem(ctx context.Context, it *models.Item) (*models.Item, error)
	UpdateItem(ctx context.Context, it *models.Item) (*models.Item, error)
	DeleteItem(ctx context.Context, id uuid.UUID) error
	Fetch(ctx context.Context, page int) (*[]models.Item, error)
}
