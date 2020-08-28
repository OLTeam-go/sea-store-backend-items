package usecase

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-items/models"
	"github.com/google/uuid"
)

// Usecase are the business logic used for this service
type Usecase interface {
	StoreItem(ctx context.Context, it *models.Item) (*models.Item, error)
	DeleteItem(id uuid.UUID) error
	UpdateItem(id uuid.UUID, it *models.Item) error
	GetByID(id uuid.UUID) (*models.Item, error)
	GetByMerchantID(merchantID uuid.UUID) ([]*models.Item, error)
}
