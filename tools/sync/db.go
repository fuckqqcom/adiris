package main

import (
	"adiris/model/admin"
	"adiris/pkg/config"
)

func main() {
	config.InitConfig("config/config.json")
	config.EngDb.Sync2(new(admin.User))
	config.EngDb.Sync2(new(admin.Group))
	config.EngDb.Sync2(new(admin.Log))
	config.EngDb.Sync2(new(admin.Menu))
	config.EngDb.Sync2(new(admin.Role))
	config.EngDb.Sync2(new(admin.RoleGroup))
	config.EngDb.Sync2(new(admin.RoleMenu))
	config.EngDb.Sync2(new(admin.UserRole))
	config.EngDb.Sync2(new(admin.UserMenu))
	config.EngDb.Sync2(new(admin.UserGroup))

}
