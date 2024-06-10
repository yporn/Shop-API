package repository

import "github.com/yporn/shop-go-api/entities"

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
}