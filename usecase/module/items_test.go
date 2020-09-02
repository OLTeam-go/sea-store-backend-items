package module_test

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/OLTeam-go/sea-store-backend-items/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mocks "github.com/OLTeam-go/sea-store-backend-items/mocks/repository"
	uItems "github.com/OLTeam-go/sea-store-backend-items/usecase/module"
)

func TestFetch(t *testing.T) {
	mockRepo := new(mocks.Repository)
	randId1, _ := uuid.NewRandom()
	randId2, _ := uuid.NewRandom()
	mockItem := models.Item{
		ID:          randId1,
		MerchantID:  randId2,
		Name:        "Item 1",
		Description: "Description 1",
		Price:       100000,
		Category:    "Category 1",
		Quantity:    1,
		CreatedAt:   time.Now(),
	}

	var mockListItem []models.Item
	var emptyListItem []models.Item
	mockListItem = append(mockListItem, mockItem)

	mockRepo.On("Fetch", mock.Anything, 1).Return(&mockListItem, nil)
	mockRepo.On("Fetch", mock.Anything, 0).Return(nil, errors.New("page is invalid"))
	mockRepo.On("Fetch", mock.Anything, 2).Return(&emptyListItem, nil)

	ucase := uItems.New(mockRepo, time.Second*2)

	t.Run("success on page 1", func(t *testing.T) {
		fetchResult, err := ucase.Fetch(context.TODO(), 1)

		assert.Len(t, *fetchResult, 1)
		assert.NoError(t, err)
	})

	t.Run("success return empty when page over limit", func(t *testing.T) {
		fetchResult, err := ucase.Fetch(context.TODO(), 2)

		assert.Len(t, *fetchResult, 0)
		assert.NoError(t, err)
	})

	t.Run("error when page is invalid", func(t *testing.T) {
		_, err := ucase.Fetch(context.TODO(), 0)

		assert.Error(t, err)
	})
}

func TestGetByMerchantID(t *testing.T) {
	mockRepo := new(mocks.Repository)
	ItemID1, _ := uuid.NewRandom()
	MerchantID1, _ := uuid.NewRandom()
	MerchantID2, _ := uuid.NewRandom()
	mockItem := models.Item{
		ID:          ItemID1,
		MerchantID:  MerchantID1,
		Name:        "Item 1",
		Description: "Description 1",
		Price:       100000,
		Category:    "Category 1",
		Quantity:    1,
		CreatedAt:   time.Now(),
	}

	var mockListItem []models.Item
	var emptyListItem []models.Item
	mockListItem = append(mockListItem, mockItem)

	mockRepo.On("GetByMerchantID", mock.Anything, MerchantID1.String(), 1).Return(&mockListItem, nil)
	mockRepo.On("GetByMerchantID", mock.Anything, MerchantID2.String(), 1).Return(&emptyListItem, nil)
	mockRepo.On("GetByMerchantID", mock.Anything, mock.AnythingOfType("string"), 0).Return(nil, errors.New("page is invalid"))
	mockRepo.On("GetByMerchantID", mock.Anything, mock.AnythingOfType("string"), 2).Return(&emptyListItem, nil)

	ucase := uItems.New(mockRepo, time.Second*2)

	t.Run("error when page invalid", func(t *testing.T) {
		fetchResult, err := ucase.GetByMerchantID(context.TODO(), MerchantID1.String(), 0)

		assert.Nil(t, fetchResult)
		assert.Error(t, err)
	})

	t.Run("success when id exists on page 1", func(t *testing.T) {
		fetchResult, err := ucase.GetByMerchantID(context.TODO(), MerchantID1.String(), 1)

		assert.Len(t, *fetchResult, 1)
		assert.Equal(t, (*fetchResult)[0].MerchantID, MerchantID1)
		assert.NoError(t, err)
	})

	t.Run("success when id not exists", func(t *testing.T) {
		fetchResult, err := ucase.GetByMerchantID(context.TODO(), MerchantID2.String(), 1)

		assert.Len(t, *fetchResult, 0)
		assert.NoError(t, err)
	})

}

func TestGetByID(t *testing.T) {
	mockRepo := new(mocks.Repository)
	ItemID1, _ := uuid.NewRandom()
	MerchantID1, _ := uuid.NewRandom()
	mockItem := models.Item{
		ID:          ItemID1,
		MerchantID:  MerchantID1,
		Name:        "Item 1",
		Description: "Description 1",
		Price:       100000,
		Category:    "Category 1",
		Quantity:    1,
		CreatedAt:   time.Now(),
	}

	mockRepo.On("GetByID", mock.Anything, ItemID1.String()).Return(&mockItem, nil)
	mockRepo.On("GetByID", mock.Anything, mock.AnythingOfType("string")).Return(nil, models.ErrNotFound)

	ucase := uItems.New(mockRepo, time.Second*2)
	t.Run("success when id matches", func(t *testing.T) {
		fetchResult, err := ucase.GetByID(context.TODO(), ItemID1.String())

		assert.Equal(t, *fetchResult, mockItem)
		assert.NoError(t, err)
	})

	t.Run("error when id not matches", func(t *testing.T) {
		_, err := ucase.GetByID(context.TODO(), "123456")

		assert.Error(t, err)
		assert.Equal(t, err, models.ErrNotFound)
	})
}

func TestDeleteItem(t *testing.T) {
	mockRepo := new(mocks.Repository)
	ItemID1, _ := uuid.NewRandom()
	ItemID2, _ := uuid.NewRandom()
	MerchantID1, _ := uuid.NewRandom()
	mockItem := models.Item{
		ID:          ItemID1,
		MerchantID:  MerchantID1,
		Name:        "Item 1",
		Description: "Description 1",
		Price:       100000,
		Category:    "Category 1",
		Quantity:    1,
		CreatedAt:   time.Now(),
	}
	mockDeletedItem := mockItem
	mockDeletedItem.DeletedAt = time.Now()

	mockRepo.On("DeleteItem", mock.Anything, ItemID1.String()).Return(&mockDeletedItem, nil)
	mockRepo.On("DeleteItem", mock.Anything, mock.AnythingOfType("string")).Return(nil, models.ErrNotFound)

	ucase := uItems.New(mockRepo, time.Second*2)
	t.Run("success delete when id matches", func(t *testing.T) {
		fetchResult, err := ucase.DeleteItem(context.TODO(), ItemID1.String())
		fmt.Println(fetchResult.DeletedAt)
		assert.NotEqual(t, (*fetchResult).DeletedAt, time.Time{})
		assert.NoError(t, err)
	})

	t.Run("error delete when id not exists", func(t *testing.T) {
		_, err := ucase.DeleteItem(context.TODO(), ItemID2.String())

		assert.Error(t, err)
		assert.Equal(t, err, models.ErrNotFound)
	})

}

func TestUpdateItem(t *testing.T) {
	mockRepo := new(mocks.Repository)
	ItemID1, _ := uuid.NewRandom()
	MerchantID1, _ := uuid.NewRandom()
	mockItem := models.Item{
		ID:          ItemID1,
		MerchantID:  MerchantID1,
		Name:        "Item 1",
		Description: "Description 1",
		Price:       50000,
		Category:    "Category 1",
		Quantity:    1,
		CreatedAt:   time.Now(),
	}
	mockUpdatedItem := models.Item{
		ID:          ItemID1,
		MerchantID:  MerchantID1,
		Name:        "update 1",
		Description: "Description 1",
		Price:       100000,
		Category:    "Category 1",
		Quantity:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("success when id matches and item valid", func(t *testing.T) {
		mockRepo.On("UpdateItem", mock.Anything, ItemID1.String(), &mockItem).Return(&mockUpdatedItem, nil)
		ucase := uItems.New(mockRepo, time.Second*2)
		fetchResult, err := ucase.UpdateItem(context.TODO(), ItemID1.String(), &mockItem)

		assert.NotEqual(t, (*fetchResult).UpdatedAt, mockItem.UpdatedAt)
		assert.NotEqual(t, (*fetchResult).Price, mockItem.Price)
		assert.NoError(t, err)
	})

	t.Run("success not update id when id did not matches", func(t *testing.T) {
		newID, _ := uuid.NewRandom()
		mockRepo.On("UpdateItem", mock.Anything, mock.AnythingOfType("string"), &mockItem).Return(&mockUpdatedItem, nil)
		ucase := uItems.New(mockRepo, time.Second*2)
		fetchResult, err := ucase.UpdateItem(context.TODO(), newID.String(), &mockItem)

		assert.NotEqual(t, (*fetchResult).UpdatedAt, mockItem.UpdatedAt)
		assert.NotEqual(t, (*fetchResult).Price, mockItem.Price)
		assert.Equal(t, (*fetchResult).ID, mockItem.ID)
		assert.NoError(t, err)
	})

	t.Run("error update price when price is below 1", func(t *testing.T) {
		newID, _ := uuid.NewRandom()
		mockInvalidItem := models.Item{
			ID:          ItemID1,
			MerchantID:  MerchantID1,
			Name:        "Item 1",
			Description: "Description 1",
			Price:       -1,
			Category:    "Category 1",
			Quantity:    1,
			CreatedAt:   time.Now(),
		}
		mockRepo.On("UpdateItem", mock.Anything, mock.AnythingOfType("string"), &mockInvalidItem).Return(nil, models.ErrBadParamInput)
		ucase := uItems.New(mockRepo, time.Second*2)
		fetchResult, err := ucase.UpdateItem(context.TODO(), newID.String(), &mockInvalidItem)

		assert.Nil(t, fetchResult)
		assert.Error(t, err)
	})

	t.Run("error update quantity when quantity is negative", func(t *testing.T) {
		newID, _ := uuid.NewRandom()
		mockInvalidItem := models.Item{
			ID:          ItemID1,
			MerchantID:  MerchantID1,
			Name:        "Item 1",
			Description: "Description 1",
			Price:       10000,
			Category:    "Category 1",
			Quantity:    -1,
			CreatedAt:   time.Now(),
		}
		mockRepo.On("UpdateItem", mock.Anything, mock.AnythingOfType("string"), &mockInvalidItem).Return(nil, models.ErrBadParamInput)
		ucase := uItems.New(mockRepo, time.Second*2)
		fetchResult, err := ucase.UpdateItem(context.TODO(), newID.String(), &mockInvalidItem)

		assert.Nil(t, fetchResult)
		assert.Error(t, err)
	})
}

func TestStoreItem(t *testing.T) {
	mockRepo := new(mocks.Repository)
	ItemID1, _ := uuid.NewRandom()
	MerchantID1, _ := uuid.NewRandom()
	mockItem := models.Item{
		ID:          ItemID1,
		MerchantID:  MerchantID1,
		Name:        "Item 1",
		Description: "Description 1",
		Price:       50000,
		Category:    "Category 1",
		Quantity:    1,
	}

	t.Run("success when id matches and item valid", func(t *testing.T) {
		mockRepo.On("StoreItem", mock.Anything, &mockItem).Return(&mockItem, nil)
		ucase := uItems.New(mockRepo, time.Second*2)
		fetchResult, err := ucase.StoreItem(context.TODO(), &mockItem)

		assert.True(t, reflect.DeepEqual(*fetchResult, mockItem))
		assert.NoError(t, err)
	})

	t.Run("error when item's price is below 1", func(t *testing.T) {
		mockItem := models.Item{
			ID:          ItemID1,
			MerchantID:  MerchantID1,
			Name:        "Item 1",
			Description: "Description 1",
			Price:       0,
			Category:    "Category 1",
			Quantity:    1,
		}
		mockRepo.On("StoreItem", mock.Anything, &mockItem).Return(nil, models.ErrBadParamInput)
		ucase := uItems.New(mockRepo, time.Second*2)
		fetchResult, err := ucase.StoreItem(context.TODO(), &mockItem)

		assert.Nil(t, fetchResult)
		assert.Error(t, err)
		assert.Equal(t, err, models.ErrBadParamInput)
	})

}
