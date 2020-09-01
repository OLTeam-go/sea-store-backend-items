package module

import (
	"context"

	"github.com/OLTeam-go/sea-store-backend-items/models"
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

func (u *itemUsecase) UpdateItem(c context.Context, id string, it *models.Item) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	res, err := u.repo.UpdateItem(ctx, id, it)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (u *itemUsecase) DeleteItem(c context.Context, id string) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	item, err := u.repo.DeleteItem(ctx, id)

	return item, err
}

func (u *itemUsecase) GetByID(c context.Context, id string) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	res, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *itemUsecase) GetByMerchantID(c context.Context, merchantID string, page int) (*[]models.Item, error) {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	res, err := u.repo.GetByMerchantID(ctx, merchantID, page)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (u *itemUsecase) Fetch(c context.Context, page int) (*[]models.Item, error) {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	res, err := u.repo.Fetch(ctx, page)

	if err != nil {
		return nil, err
	}

	return res, err
}
