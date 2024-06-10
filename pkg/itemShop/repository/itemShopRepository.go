package repository

import (
	"github.com/yporn/shop-go-api/entities"
	_itemShopModel "github.com/yporn/shop-go-api/pkg/itemShop/model"

)

type ItemShopRepository interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
	Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error)
}
