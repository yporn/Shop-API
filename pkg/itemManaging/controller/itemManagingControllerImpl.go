package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yporn/shop-go-api/pkg/custom"
	_itemManagingModel "github.com/yporn/shop-go-api/pkg/itemManaging/model"
	_itemManagingService "github.com/yporn/shop-go-api/pkg/itemManaging/service"
)

type itemManagingControllerImpl struct {
	itemManagingService _itemManagingService.ItemManagingService
}

func NewItemManagingController(itemManagingService _itemManagingService.ItemManagingService) ItemManagingController {
	return &itemManagingControllerImpl{itemManagingService}
}

func (c *itemManagingControllerImpl) Creating(pctx echo.Context) error {
	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)
	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemCreatingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	item, err := c.itemManagingService.Creating(itemCreatingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusCreated, item)
}
