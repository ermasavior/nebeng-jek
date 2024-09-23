package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"nebeng-jek/internal/modules/products/model"
	"nebeng-jek/pkg/utils"
)

func (usecase *productUsecase) GetProductByID(ctx context.Context, id string) (*model.Product, error) {
	productCache, _ := usecase.redis.Get(ctx, fmt.Sprintf("%s:%s", utils.RedisKeyGetProductDetail, id)).Result()
	if productCache != "" {
		var productData model.Product
		if err := json.Unmarshal([]byte(productCache), &productData); err != nil {
			return nil, err
		}
		return &productData, nil
	}
	product, err := usecase.repository.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	marshaleProductData, _ := json.Marshal(product)
	usecase.redis.Set(ctx, fmt.Sprintf("%s:%s", utils.RedisKeyGetProductDetail, id), marshaleProductData, 3*time.Minute)
	return &product, nil
}
