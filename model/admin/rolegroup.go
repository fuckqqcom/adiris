package admin

import (
	"adiris/pkg/config"
	"adiris/tools/common"
)

type RoleGroup struct {
	Id     string //主键id
	IsDel  int
	Remark string
	Status int
	Rid    string
	Gid    string
	at     `xorm:"extends"`
}

/**
给用户分配对应的角色
一个用户只能对应一个角色
*/

func AddRoleGroup(rid, gid, remark string) int {
	r := RoleGroup{Id: commons.EncodeMd5(commons.StringJoin(rid, gid)), IsDel: 1, Remark: remark, Status: 1, Rid: rid, Gid: gid}
	return CheckInt64(config.EngDb.Insert(r))
}

/**
删除用户角色关联关系
*/

func DeleteRoleGroup(id string) int {
	r := RoleGroup{Id: id}
	return CheckInt64(config.EngDb.Where("id = ? ", id).Delete(r))
}

/**
修改用户角色关联关系
*/

func UpdateRoleGroup(rid, gid, remark string) int {
	r := RoleGroup{Id: commons.EncodeMd5(commons.StringJoin(rid, gid)), Rid: rid, Remark: remark, Gid: gid}
	return CheckInt64(config.EngDb.Where("id = ? ", r.Id).Cols("remark", "rid", "gid").Update(r))

}
