package rest

import (
	"net/http"
	"time"

	"github.com/OLTeam-go/sea-store-backend-items/models"
	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
)

// ResponseError represents response when error occurs
type ResponseError struct {
	Message   string    `json:"message"`
	timestamp time.Time `default:"time.Now()"`
}

// ResponseSuccess represents response when success
type ResponseSuccess struct {
	Status    int         `json:"status"`
	Data      interface{} `json:"data"`
	timestamp time.Time   `default:"time.Now()"`
}

func isRequestValid(m *models.Item) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *restDelivery) StoreItem(c echo.Context) error {
	var item models.Item
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseError{
			Message: err.Error(),
		})
	}
	if ok, err := isRequestValid(&item); !ok {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
	}
	ctx := c.Request().Context()
	res, err := r.usecase.StoreItem(ctx, &item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, ResponseSuccess{
		Status: http.StatusCreated,
		Data:   res,
	})
}

func (r *restDelivery) UpdateItem() {

}

func (r *restDelivery) DeleteItem() {

}

func (r *restDelivery) GetByID() {

}

func (r *restDelivery) GetByMerchantID() {

}

func (r *restDelivery) Test(c echo.Context) error {
	return c.JSON(http.StatusOK, "oi")
}
