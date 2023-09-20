package main

import (
	"fmt"

	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/cache"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
)

func main() {
	cfg := config.GetConfig()
	fmt.Println(cfg)
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

}
