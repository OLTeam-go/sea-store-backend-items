package rest

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/OLTeam-go/sea-store-backend-items/models"
	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
)

// ResponsePagination represent the response for pagination request
type ResponsePagination struct {
	Status int         `json:"status"`
	Page   int         `json:"page"`
	Size   int         `json:"size"`
	Data   interface{} `json:"data"`
}

// ResponseError represents response when error occurs
type ResponseError struct {
	Message string `json:"message"`
}

// ResponseSuccess represents response when success
type ResponseSuccess struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func isRequestValid(m *models.Item) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func parsePagination(c echo.Context) (int, error) {
	pageQuery := c.QueryParam("page")
	if pageQuery == "" {
		pageQuery = "1"
	}
	page, err := strconv.Atoi(pageQuery)
	if page <= 0 {
		return 0, errors.New("page is invalid")
	}
	return page, err
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

func (r *restDelivery) UpdateItem(c echo.Context) error {
	var item models.Item
	id := c.Param("id")
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
	res, err := r.usecase.UpdateItem(ctx, id, &item)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ResponseSuccess{
		Status: http.StatusAccepted,
		Data:   res,
	})
}

func (r *restDelivery) DeleteItem(c echo.Context) error {
	var item models.Item
	id := c.Param("id")

	ctx := c.Request().Context()
	err := r.usecase.DeleteItem(ctx, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, ResponseSuccess{
		Status: http.StatusAccepted,
		Data:   item.ID,
	})
}

func (r *restDelivery) GetByID(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	res, err := r.usecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ResponseSuccess{
		Status: http.StatusOK,
		Data:   res,
	})
}

func (r *restDelivery) Fetch(c echo.Context) error {
	page, err := parsePagination(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	res, err := r.usecase.Fetch(ctx, page)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, ResponsePagination{
		Status: http.StatusAccepted,
		Page:   page,
		Data:   res,
		Size:   len(*res),
	})

}

func (r *restDelivery) GetByMerchantID(c echo.Context) error {
	merchantID := c.Param("merchant_id")
	page, err := parsePagination(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	res, err := r.usecase.GetByMerchantID(ctx, merchantID, page)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ResponsePagination{
		Status: http.StatusOK,
		Data:   res,
		Page:   page,
		Size:   len(*res),
	})

}
