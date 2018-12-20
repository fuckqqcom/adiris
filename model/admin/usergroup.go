package admin

import (
	"adiris/pkg/config"
	"adiris/tools/common"
)

type UserGroup struct {
	Id     string `xorm:"varchar(100) notnull pk index unique 'id'"` //主键id
	IsDel  int
	Remark string
	Status int
	Uid    string
	Gid    string
	at     `xorm:"extends"`
}

/**
给用户分配对应的角色
一个用户只能对应一个角色
*/

func AddUserGroup(uid, gid, remark string) int {
	u := UserGroup{Id: commons.EncodeMd5(commons.StringJoin(uid, gid)), IsDel: 1, Remark: remark, Status: 1, Uid: uid, Gid: gid}
	return CheckInt64(config.EngDb.Insert(u))
}

/**
删除用户角色关联关系
*/

func DeleteUserGroup(id string) int {
	u := UserGroup{Id: id}
	return CheckInt64(config.EngDb.Where("id = ? ", id).Delete(u))
}

/**
修改用户角色关联关系
*/

func UpdateUserGroup(uid, gid, remark string) int {
	r := UserGroup{Id: commons.EncodeMd5(commons.StringJoin(uid, gid)), Uid: uid, Remark: remark, Gid: gid}
	return CheckInt64(config.EngDb.Where("id = ? ", r.Id).Cols("remark", "uid", "gid").Update(r))

}
