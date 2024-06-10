package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/yporn/shop-go-api/entities"
	"gorm.io/gorm"

	_itemShopExeption "github.com/yporn/shop-go-api/pkg/itemManaging/exception"
)

type itemManagingRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemManagingRepository {
	return &itemManagingRepositoryImpl{db, logger}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Errorf("Creating item failed: %s", err.Error())
		return nil, &_itemShopExeption.ItemCreating{}
	}

	return item, nil
}
