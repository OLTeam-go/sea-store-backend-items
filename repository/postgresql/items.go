package postgresql

import (
	"context"

	"github.com/google/uuid"

	"github.com/OLTeam-go/sea-store-backend-items/models"
)

func (r *postgresqlRepository) StoreItem(ctx context.Context, it *models.Item) (*models.Item, error) {
	var item models.Item
	item.MerchantID = it.MerchantID
	item.Name = it.Name
	item.Description = it.Description
	item.Price = it.Price
	item.Quantity = it.Quantity
	item.Category = it.Category

	_, err := r.Conn.Model(&item).Returning("*").Insert()

	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *postgresqlRepository) UpdateItem(ctx context.Context, it *models.Item) error {
	return nil
}

func (r *postgresqlRepository) DeleteItem(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (r *postgresqlRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Item, error) {
	return nil, nil
}

func (r *postgresqlRepository) GetByMerchantID(ctx context.Context, id uuid.UUID) ([]*models.Item, error) {
	return nil, nil
}
