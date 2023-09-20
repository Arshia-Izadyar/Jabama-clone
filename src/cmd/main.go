package main

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/api"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/cache"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db/migrations"
)

func main() {
	cfg := config.GetConfig()
	err := db.InitDB(cfg)
	defer db.CloseDB()
	if err != nil {
		panic(err)
	}
	err = cache.InitRedis(cfg)
	if err != nil {
		panic(err)
	}
	defer cache.CloseRedis()
	migrations.Up_01()
	api.Init(cfg)

}
