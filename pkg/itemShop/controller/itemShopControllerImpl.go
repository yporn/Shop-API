package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yporn/shop-go-api/pkg/custom"
	// _itemShopExeption "github.com/yporn/shop-go-api/pkg/itemShop/exception"
	_itemShopModel "github.com/yporn/shop-go-api/pkg/itemShop/model"
	_itemShopService "github.com/yporn/shop-go-api/pkg/itemShop/service"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemFilter := new(_itemShopModel.ItemFilter)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemFilter); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	itemModelList, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, itemModelList)
	// return custom.CustomError(pctx, http.StatusInternalServerError, (&_itemShopExeption.ItemListing{}).Error())

}
