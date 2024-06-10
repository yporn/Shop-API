package main

import (
	"github.com/yporn/shop-go-api/config"
	"github.com/yporn/shop-go-api/databases"
	"github.com/yporn/shop-go-api/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.ConnectionGetting())

	server.Start()
}