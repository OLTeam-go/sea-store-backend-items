package postgresql

import (
	"context"
	"log"
	"time"

	"github.com/OLTeam-go/sea-store-backend-items/models"
	"github.com/google/uuid"
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

func (r *postgresqlRepository) UpdateItem(ctx context.Context, id string, it *models.Item) (*models.Item, error) {
	var item models.Item
	item.Name = it.Name
	item.Category = it.Category
	item.Price = it.Price
	item.Quantity = it.Quantity
	item.Description = it.Description
	item.UpdatedAt = time.Now()

	_, err := r.Conn.Model(&item).
		Column("name", "category", "description", "price", "quantity", "updated_at").
		Where("id = ?", id).
		Returning("*").
		UpdateNotNull(&item)

	if err != nil {
		return nil, err
	}

	return &item, err
}

func (r *postgresqlRepository) DeleteItem(ctx context.Context, id string) (*models.Item, error) {
	var item models.Item
	now := time.Now()
	_, err := r.Conn.Model(&item).
		Set("deleted_at = ?", now).
		Where("id = ? AND deleted_at is NULL", id).
		Returning("*").
		Update()

	return &item, err
}

func (r *postgresqlRepository) GetByID(ctx context.Context, id string) (*models.Item, error) {
	var item models.Item
	err := r.Conn.Model(&item).
		Where("id = ? AND deleted_at is NULL", id).
		Limit(1).
		Select()

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *postgresqlRepository) GetByMerchantID(ctx context.Context, merchantID string, page int) (*[]models.Item, error) {
	var items []models.Item
	var offset int
	offset = (page - 1) * r.pagesize
	limit := r.pagesize
	var err error
	if page != 0 {
		err = r.Conn.Model(&items).
			Where("merchant_id = ? AND deleted_at is NULL", merchantID).
			Offset(offset).
			Limit(limit).
			Returning("*").
			Select()
	} else {
		err = r.Conn.Model(&items).
			Where("merchant_id = ? AND deleted_at is NULL", merchantID).
			Returning("*").
			Select()
	}

	if err != nil {
		return nil, err
	}

	return &items, err
}

func (r *postgresqlRepository) Fetch(ctx context.Context, page int) (*[]models.Item, error) {
	var items []models.Item
	var offset int
	offset = (page - 1) * r.pagesize
	limit := r.pagesize
	var err error
	if page != 0 {
		err = r.Conn.Model(&items).
			Where("deleted_at is NULL").
			Order("created_at ASC").
			Offset(offset).
			Limit(limit).
			Returning("*").
			Select()
	} else {
		log.Println("all")
		err = r.Conn.Model(&items).
			Where("deleted_at is NULL").
			Order("created_at ASC").
			Returning("*").
			Select()
	}

	if err != nil {
		return nil, err
	}

	return &items, err
}

func (r *postgresqlRepository) FetchByIDs(ctx context.Context, ids []uuid.UUID) (*[]models.Item, error) {
	var items []models.Item
	err := r.Conn.Model(&items).WhereIn("id IN (?)", ids).Returning("*").Select()
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (r *postgresqlRepository) Sold(ctx context.Context, ids []uuid.UUID) error {
	var items []models.Item
	err := r.Conn.Model(&items).
		WhereIn("id IN (?)", ids).
		Returning("*").Select()

	for _, item := range items {
		log.Println(item)
		if item.Quantity != 0 {
			item.UpdatedAt = time.Now()
			item.Quantity = item.Quantity - 1
			_, err = r.Conn.Model(&item).
				Column("quantity", "updated_at").
				Where("id = ?", item.ID).
				Returning("*").
				Update()
			log.Println(err)
			log.Println(item)
		}
	}
	return err
}

func (r *postgresqlRepository) SetAvailable(ctx context.Context, ids []uuid.UUID) error {
	var items []models.Item
	err := r.Conn.Model(&items).
		WhereIn("id IN (?)", ids).
		Returning("*").Select()

	for _, item := range items {
		if item.Quantity != 0 {
			item.UpdatedAt = time.Now()
			item.Quantity = item.Quantity + 1
			_, _ = r.Conn.Model(&item).
				Column("quantity", "updated_at").
				Where("id = ?", item.ID).
				Returning("*").
				Update()
		}
	}
	return err
}
