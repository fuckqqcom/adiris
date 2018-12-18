package admin

import (
	"adiris/pkg/config"
	"adiris/pkg/e"
	"adiris/tools/common"
)

type Role struct {
	Id     string //主键id
	IsDel  int    // -1表示删除
	Remark string
	Status int
	Rid    string
	Name   string
	at     `xorm:"extends"`
}

/**
角色的增删改查
*/
func (r Role) AddRole() int {
	if GetIdByRole(r.Id) {
		return e.RoleExist
	}
	return CheckInt64(config.EngDb.Insert(r))
}

func GetIdByRole(id string) bool {
	return CheckBool(config.EngDb.Where("id = ? and status = 1 and is_del = 1 ", id).Exist(&Role{}))
}

func CheckBool(exist bool, err error) bool {
	if commons.CheckErr(err, exist) && exist {
		return true
	} else {
		return false
	}
}

func CheckInt64(exist int64, err error) int {
	if commons.CheckErr(err, exist) && exist != 0 {
		return e.Success
	} else {
		return e.Error
	}
}
