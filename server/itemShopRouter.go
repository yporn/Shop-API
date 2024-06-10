package server

import (
	_itemShopRepository "github.com/yporn/shop-go-api/pkg/itemShop/repository"
	_itemShopService "github.com/yporn/shop-go-api/pkg/itemShop/service"
	_itemShopController "github.com/yporn/shop-go-api/pkg/itemShop/controller"
)

func (s *echoServer) initItemShopRouter() {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository)
	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("", itemShopController.Listing)
}