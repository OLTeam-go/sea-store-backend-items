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
	e.GET("/test", handler.Test)
	e.POST("/item/add", handler.StoreItem)
	return handler
}
