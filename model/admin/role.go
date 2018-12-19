package admin

import (
	"adiris/pkg/config"
	"adiris/pkg/e"
	"adiris/tools/common"
)

type Role struct {
	Id     string `xorm:"id" json:"id"`         //主键id
	IsDel  int    `xorm:"is_del" json:"is_del"` // -1表示删除
	Remark string `xorm:"remark" json:"remark"`
	Status int    `xorm:"status" json:"status"`
	Name   string `xorm:"name" json:"name"`
	at     `xorm:"extends"`
}

/**
角色的增删改查
*/
func AddRole(name, remark string, status int) int {
	r := Role{Id: commons.EncodeMd5(name), IsDel: 1, Name: name, Status: status, Remark: remark}

	if GetIdByRole(r.Id) {
		return e.RoleExist
	}
	return CheckInt64(config.EngDb.Insert(r))
}

func GetIdByRole(id string) bool {
	return CheckBool(config.EngDb.Where("id = ? and status = 1 and is_del = 1 ", id).Exist(&Role{}))
}

/**
删除角色:单个删除
*/
func DeleteRole(id string) int {
	r := Role{Id: id}
	return CheckInt64(config.EngDb.Where("id = ?", id).Delete(r))
}

/**
修改角色
*/

func UpdateRole(name, remark string, status int) int {
	r := Role{Id: commons.EncodeMd5(name), IsDel: 1, Name: name, Status: status, Remark: remark}
	return CheckInt64(config.EngDb.Where("id = ? ", r.Id).Cols("remark", "status", "name").Update(r))
}

/**
查询角色集合
*/

func GetRoleList(pn, ps int) interface{} {
	var r []Role
	count, err := config.EngDb.Where("is_del = 1 and status = 1").Desc("create_time").Limit(ps, (pn-1)*ps).FindAndCount(&r)
	CheckInt64(count, err)
	m := make(map[string]interface{})
	m["count"] = int(count)
	m["data"] = r
	return m
}
