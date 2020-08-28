package rest

import (
	dItems "github.com/OLTeam-go/sea-store-backend-items/delivery"
	uItems "github.com/OLTeam-go/sea-store-backend-items/usecase"
	"github.com/labstack/echo"
)

type restDelivery struct {
	usecase uItems.Usecase
}

// New function initialize delivery used for the services
func New(e *echo.Echo, usecase uItems.Usecase) dItems.Delivery {
	handler := &restDelivery{
		usecase: usecase,
	}
	e.POST("/item", handler.StoreItem)
	e.DELETE("/item", handler.DeleteItem)
	e.GET("/item/:id", handler.GetByID)
	e.PATCH("/item", handler.UpdateItem)
	e.GET("/items/merchant/:merchant_id", handler.GetByMerchantID)
	e.GET("/items", handler.Fetch)
	return handler
}
