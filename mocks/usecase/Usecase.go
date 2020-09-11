// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/OLTeam-go/sea-store-backend-items/models"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// DeleteItem provides a mock function with given fields: c, id
func (_m *Usecase) DeleteItem(c context.Context, id string) (*models.Item, error) {
	ret := _m.Called(c, id)

	var r0 *models.Item
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Item); ok {
		r0 = rf(c, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: c, page
func (_m *Usecase) Fetch(c context.Context, page int) (*[]models.Item, error) {
	ret := _m.Called(c, page)

	var r0 *[]models.Item
	if rf, ok := ret.Get(0).(func(context.Context, int) *[]models.Item); ok {
		r0 = rf(c, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(c, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchByIDs provides a mock function with given fields: c, id
func (_m *Usecase) FetchByIDs(c context.Context, id []uuid.UUID) (*[]models.Item, error) {
	ret := _m.Called(c, id)

	var r0 *[]models.Item
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) *[]models.Item); ok {
		r0 = rf(c, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []uuid.UUID) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: c, id
func (_m *Usecase) GetByID(c context.Context, id string) (*models.Item, error) {
	ret := _m.Called(c, id)

	var r0 *models.Item
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Item); ok {
		r0 = rf(c, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByMerchantID provides a mock function with given fields: c, merchantID, page
func (_m *Usecase) GetByMerchantID(c context.Context, merchantID string, page int) (*[]models.Item, error) {
	ret := _m.Called(c, merchantID, page)

	var r0 *[]models.Item
	if rf, ok := ret.Get(0).(func(context.Context, string, int) *[]models.Item); ok {
		r0 = rf(c, merchantID, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(c, merchantID, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sold provides a mock function with given fields: c, id
func (_m *Usecase) Sold(c context.Context, id []uuid.UUID) error {
	ret := _m.Called(c, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) error); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreItem provides a mock function with given fields: c, it
func (_m *Usecase) StoreItem(c context.Context, it *models.Item) (*models.Item, error) {
	ret := _m.Called(c, it)

	var r0 *models.Item
	if rf, ok := ret.Get(0).(func(context.Context, *models.Item) *models.Item); ok {
		r0 = rf(c, it)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Item) error); ok {
		r1 = rf(c, it)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateItem provides a mock function with given fields: c, id, it
func (_m *Usecase) UpdateItem(c context.Context, id string, it *models.Item) (*models.Item, error) {
	ret := _m.Called(c, id, it)

	var r0 *models.Item
	if rf, ok := ret.Get(0).(func(context.Context, string, *models.Item) *models.Item); ok {
		r0 = rf(c, id, it)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *models.Item) error); ok {
		r1 = rf(c, id, it)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
