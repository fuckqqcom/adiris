package main

import (
	"adiris/monitor"
	"adiris/pkg/config"
)

func main() {
	//routers.InitRouter()
	config.InitConfig("config/config.json")
	monitor.RedisInfo()
}
