package main

import (
	"adiris/model/admin"
	"adiris/pkg/config"
)

func main() {
	config.InitConfig("config/config.json")
	config.EngDb.Sync2(new(admin.User))
	config.EngDb.Sync2(new(admin.Dept))
	config.EngDb.Sync2(new(admin.Log))
	config.EngDb.Sync2(new(admin.Menu))
	config.EngDb.Sync2(new(admin.Role))
	config.EngDb.Sync2(new(admin.RoleDept))
	config.EngDb.Sync2(new(admin.RoleMenu))
	config.EngDb.Sync2(new(admin.UserRole))
	config.EngDb.Sync2(new(admin.UserMenu))

}
