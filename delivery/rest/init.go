package rest

import (
	"net/http"

	dItems "github.com/OLTeam-go/sea-store-backend-items/delivery"
	uItems "github.com/OLTeam-go/sea-store-backend-items/usecase"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type restDelivery struct {
	usecase uItems.Usecase
}

// New function initialize delivery used for the services
func New(e *echo.Echo, usecase uItems.Usecase) dItems.Delivery {
	handler := &restDelivery{
		usecase: usecase,
	}
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusAccepted, "Hello from service items")
	})
	e.POST("/item", handler.StoreItem)
	e.DELETE("/item:id", handler.DeleteItem)
	e.GET("/item/:id", handler.GetByID)
	e.PATCH("/item/:id", handler.UpdateItem)
	e.PUT("/item/:id", handler.UpdateItem)
	e.GET("/items/merchant/:merchant_id", handler.GetByMerchantID)
	e.GET("/items", handler.Fetch)
	e.GET("/docs/*", echoSwagger.WrapHandler)
	return handler
}
