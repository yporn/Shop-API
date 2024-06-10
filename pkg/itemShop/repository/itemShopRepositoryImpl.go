package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/yporn/shop-go-api/entities"
	"gorm.io/gorm"

	_itemShopExeption "github.com/yporn/shop-go-api/pkg/itemShop/exception"
	_itemShopModel "github.com/yporn/shop-go-api/pkg/itemShop/model"
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false) // select * from items

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	size := int(itemFilter.Size)

	if err := query.Offset(offset).Limit(size).Find(&itemList).Order("id asc").Error; err != nil {
		r.logger.Errorf("Failed to list items: %s", err.Error())
		return nil, &_itemShopExeption.ItemListing{}
	}

	return itemList, nil
}

func (r *itemShopRepositoryImpl) Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error) {
	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false) // select * from items

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	var count int64

	if err := query.Count(&count).Error; err != nil {
		r.logger.Errorf("Counting items failed: %s", err.Error())
		return -1, &_itemShopExeption.ItemCounting{}
	}

	return count, nil
}
