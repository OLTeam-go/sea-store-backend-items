package rest

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/OLTeam-go/sea-store-backend-items/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
)

// RequestIDs represent the request for Fetch by multiple ID
type RequestIDs struct {
	IDs []uuid.UUID `json:"ids"`
}

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

// StoreItem process request to create new item
// @Summary Endpoint to create new item
// @Description Create new item based on json on
// @Accept json
// @Produce json
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /item [post]
// @Success 201 {object} ResponseSuccess{data=models.Item}
// @Param default body models.Item true "Created at, updated at, deleted at are optional and will be ignored"
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

// UpdateItem process request to update an item
// @Summary Endpoint to update an item
// @Description Update item based on provided data
// @Accept json
// @Produce json
// @Failure 422 {object} ResponseError
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /item/{id} [patch]
// @Success 200 {object} ResponseSuccess{data=models.Item}
// @Param default body models.Item false "Only Name, Category, Description, Quantity, and Price will be updated"
// @Param id path string true "ID of an item"
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
		Status: http.StatusOK,
		Data:   res,
	})
}

// DeleteItem process request to delete an item
// @Summary Endpoint to delete an item
// @Description Delete an item based on the id
// @Accept json
// @Produce json
// @Failure 404 {object} ResponseError
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /item/{id} [delete]
// @Success 200 {object} ResponseSuccess{data=models.Item}
// @Param id path string true "ID of an item"
func (r *restDelivery) DeleteItem(c echo.Context) error {
	id := c.Param("id")

	ctx := c.Request().Context()
	item, err := r.usecase.DeleteItem(ctx, id)

	if item.ID.String() != id {
		return c.JSON(http.StatusNotFound, ResponseError{
			Message: "Item did not exists",
		})
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ResponseSuccess{
		Status: http.StatusOK,
		Data:   id,
	})
}

// GetByID process request to get an item based on its id
// @Summary Endpoint to get an item by id
// @Description return item object
// @Accept json
// @Produce json
// @Failure 404 {object} ResponseError
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /item/{id} [get]
// @Success 200 {object} ResponseSuccess{data=models.Item}
// @Param id path string true "ID of an item"
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

// Fetch process request to get items
// @Summary Endpoint to get items
// @Description return array of item object
// @Accept json
// @Produce json
// @Failure 404 {object} ResponseError
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /items [get]
// @Success 200 {object} ResponsePagination{data=[]models.Item}
// @Param page query int false "page index"
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

	return c.JSON(http.StatusOK, ResponsePagination{
		Status: http.StatusOK,
		Page:   page,
		Data:   res,
		Size:   len(*res),
	})

}

// GetByMerchantID process request to get items based on the merchant ID
// @Summary Endpoint to get items based on the merchant ID
// @Description return array of item object
// @Accept json
// @Produce json
// @Failure 404 {object} ResponseError
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /items/merchant/{merchant_id} [get]
// @Success 200 {object} ResponsePagination{data=[]models.Item}
// @Param page query int false "page index"
// @Param merchant_id path string true "merchant id"
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

// FetchByIDs process request to get items based given IDs
// @Summary Endpoint to get items based on IDs
// @Description return array of item object
// @Accept json
// @Produce json
// @Failure 404 {object} ResponseError
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /items [post]
// @Success 200 {object} ResponseSuccess
// @Param default body RequestIDs true "Request ID"
func (r *restDelivery) FetchByIDs(c echo.Context) error {
	var body RequestIDs
	c.Bind(&body)
	ctx := c.Request().Context()

	res, err := r.usecase.FetchByIDs(ctx, body.IDs)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
	}

	log.Println(res)

	return c.JSON(http.StatusOK, ResponseSuccess{
		Status: http.StatusOK,
		Data:   res,
	})
}

// Sold process request to set items to be sold
// @Summary Endpoint to set items to be sold (quantitiy = 0)
// @Accept json
// @Produce json
// @Failure 404 {object} ResponseError
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /items/sold [post]
// @Success 200 {object} ResponseSuccess
// @Param default body RequestIDs true "Request ID"
func (r *restDelivery) Sold(c echo.Context) error {
	var body RequestIDs
	c.Bind(&body)
	ctx := c.Request().Context()

	err := r.usecase.Sold(ctx, body.IDs)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ResponseSuccess{
		Status: http.StatusOK,
	})
}

// SetAvailable process request to set items to be available
// @Summary Endpoint to set items to be available (quantitiy = 1)
// @Accept json
// @Produce json
// @Failure 404 {object} ResponseError
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /items/available [post]
// @Success 200 {object} ResponseSuccess
// @Param default body RequestIDs true "Request ID"
func (r *restDelivery) SetAvailable(c echo.Context) error {
	var body RequestIDs
	c.Bind(&body)
	ctx := c.Request().Context()

	err := r.usecase.SetAvailable(ctx, body.IDs)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ResponseSuccess{
		Status: http.StatusOK,
	})
}
