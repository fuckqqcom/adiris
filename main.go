package main

import (
	"adiris/pkg/config"
	"adiris/routers"
)

func main() {
	routers.InitRouter()
	config.InitConfig("config/config.json")
	//monitor.RedisInfo()
}
