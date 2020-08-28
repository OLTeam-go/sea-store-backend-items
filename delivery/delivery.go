package delivery

import (
	"github.com/labstack/echo"
)

// Delivery represent the method used for handling request
type Delivery interface {
	StoreItem(c echo.Context) error
	DeleteItem(c echo.Context) error
	GetByID(c echo.Context) error
	UpdateItem(c echo.Context) error
	GetByMerchantID(c echo.Context) error
	Fetch(c echo.Context) error
}
