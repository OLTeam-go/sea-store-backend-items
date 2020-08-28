package module

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-items/models"
	"github.com/google/uuid"
)

func (u *itemUsecase) StoreItem(c context.Context, it *models.Item) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	dbItem, err := u.repo.StoreItem(ctx, it)
	if err != nil {
		return nil, err
	}
	return dbItem, nil
}

func (u *itemUsecase) UpdateItem(id uuid.UUID, it *models.Item) error {
	return nil
}

func (u *itemUsecase) DeleteItem(id uuid.UUID) error {
	return nil
}

func (u *itemUsecase) GetByID(id uuid.UUID) (*models.Item, error) {
	return nil, nil
}

func (u *itemUsecase) GetByMerchantID(merchantID uuid.UUID) ([]*models.Item, error) {
	return nil, nil
}
