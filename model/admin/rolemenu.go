package admin

import (
	"adiris/pkg/config"
	"adiris/tools/common"
)

type RoleMenu struct {
	Id     string //主键id
	IsDel  int
	Remark string
	Status int
	Rid    string
	Mid    string
	at     `xorm:"extends"`
}

func AddRoleMenu(rid, mid, remark string) int {
	r := RoleMenu{Id: commons.EncodeMd5(commons.StringJoin(rid, mid)), IsDel: 1, Remark: remark, Status: 1, Rid: rid, Mid: mid}
	return CheckInt64(config.EngDb.Insert(r))
}

/**
删除用户角色关联关系
*/

func DeleteRoleMenu(id string) int {
	r := RoleMenu{Id: id}
	return CheckInt64(config.EngDb.Where("id = ? ", id).Delete(r))
}

/**
修改用户角色关联关系
*/

func UpdateRoleMenu(rid, mid, remark string) int {
	r := RoleMenu{Id: commons.EncodeMd5(commons.StringJoin(rid, mid)), Rid: rid, Remark: remark, Mid: mid}
	return CheckInt64(config.EngDb.Where("id = ? ", r.Id).Cols("remark", "rid", "mid").Update(r))

}
