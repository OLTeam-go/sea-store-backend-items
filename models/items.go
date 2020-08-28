package models

import (
	"time"

	"github.com/google/uuid"
)

// Item represent the item model in the database
type Item struct {
	ID          uuid.UUID `json:"id" pg:"type:uuid,default:gen_random_uuid(),pk"`
	MerchantID  uuid.UUID `json:"merchant_id" pg:"type:uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       int64     `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at" pg:"default:now()"`
	UpdatedAt   time.Time `json:"updated_at" pg:"default:now()"`
	DeletedAt   time.Time `json:"deleted_at" pg:"soft_delete"`
}
