package main

import (
	"adiris/pkg/config"
	"adiris/routers"
)

func main() {
	config.InitConfig("config/config.json")
	routers.InitRouter()
	//monitor.RedisInfo()
}
