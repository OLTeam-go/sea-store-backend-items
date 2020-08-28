package delivery

import (
	"github.com/labstack/echo"
)

// Delivery represent the method used for handling request
type Delivery interface {
	StoreItem(e echo.Context) error
	DeleteItem()
	UpdateItem()
	GetByID()
	GetByMerchantID()
	Test(e echo.Context) error
}
