package server

import (
	_itemManagingController "github.com/yporn/shop-go-api/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/yporn/shop-go-api/pkg/itemManaging/repository"
	_itemManagingService "github.com/yporn/shop-go-api/pkg/itemManaging/service"
)

func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-managing")

	itemManagingRepository := _itemManagingRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemManagingRepository)
	itemManagingController := _itemManagingController.NewItemManagingController(itemManagingService)

	router.POST("", itemManagingController.Creating)
}
