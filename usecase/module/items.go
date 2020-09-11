package module

import (
	"context"
	"errors"

	"github.com/OLTeam-go/sea-store-backend-items/models"
	"github.com/google/uuid"
)

func validatePage(page int) (bool, error) {
	if page <= 0 {
		return false, errors.New("page is invalid")
	}
	return true, nil
}

func validateItem(it *models.Item) (bool, error) {
	if it.Quantity < 0 {
		return false, models.ErrBadParamInput
	}
	if it.Price <= 0 {
		return false, models.ErrBadParamInput
	}
	return true, nil
}

func (u *itemUsecase) StoreItem(c context.Context, it *models.Item) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	valid, err := validateItem(it)
	if !valid {
		return nil, err
	}

	dbItem, err := u.repo.StoreItem(ctx, it)
	if err != nil {
		return nil, err
	}
	return dbItem, nil
}

func (u *itemUsecase) UpdateItem(c context.Context, id string, it *models.Item) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	valid, err := validateItem(it)
	if !valid {
		return nil, err
	}

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

	if err != nil {
		return nil, err
	}

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

	valid, err := validatePage(page)

	if !valid {
		return nil, err
	}

	res, err := u.repo.GetByMerchantID(ctx, merchantID, page)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (u *itemUsecase) Fetch(c context.Context, page int) (*[]models.Item, error) {

	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	valid, err := validatePage(page)

	if !valid {
		return nil, err
	}

	res, err := u.repo.Fetch(ctx, page)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (u *itemUsecase) FetchByIDs(c context.Context, ids []uuid.UUID) (*[]models.Item, error) {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	res, err := u.repo.FetchByIDs(ctx, ids)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *itemUsecase) Sold(c context.Context, ids []uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, u.timeoutContext)
	defer cancel()

	err := u.repo.Sold(ctx, ids)
	return err
}
