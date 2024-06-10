package service

import (
	_itemManagingModel "github.com/yporn/shop-go-api/pkg/itemManaging/model"
	_itemShopModel "github.com/yporn/shop-go-api/pkg/itemShop/model"
)

type ItemManagingService interface {
	Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error)
}
