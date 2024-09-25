package usecase

import (
	"context"
	"errors"
	"testing"

	"nebeng-jek/internal/products/model"
	mockRepo "nebeng-jek/mock/repository"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mockRepo.NewMockProductRepository(ctrl)
	usecaseMock := productUsecase{
		repository: repoMock,
	}

	ctx := context.Background()
	req := model.CreateProduct{
		Name:  "Laptop Lenovo XYZ",
		Price: 30000000,
	}
	productID := "c5f08a52-cc46-47d1-879b-15e120885366"

	t.Run("success - should create new product", func(t *testing.T) {
		repoMock.EXPECT().CreateProduct(gomock.Any(), req).
			Return(productID, nil)

		id, err := usecaseMock.CreateProduct(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, productID, id)
	})

	t.Run("failed - should return error when create new product is failed", func(t *testing.T) {
		expectedErr := errors.New("db down")

		repoMock.EXPECT().CreateProduct(gomock.Any(), req).
			Return("", expectedErr)

		id, err := usecaseMock.CreateProduct(ctx, req)

		assert.Equal(t, expectedErr, err)
		assert.Equal(t, "", id)
	})
}
