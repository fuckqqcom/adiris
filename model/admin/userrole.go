package admin

import (
	"adiris/pkg/config"
	"adiris/tools/common"
)

type UserRole struct {
	Id     string `xorm:"varchar(100) notnull pk index unique 'id'"` //主键id
	IsDel  int
	Remark string
	Status int
	Uid    string
	Rid    string
	at     `xorm:"extends"`
}

/**
给用户分配对应的角色
一个用户只能对应一个角色
*/

func AddUserRole(uid, rid, remark string) int {
	u := UserRole{Id: commons.EncodeMd5(commons.StringJoin(uid, rid)), IsDel: 1, Remark: remark, Status: 1, Uid: uid, Rid: rid}
	return CheckInt64(config.EngDb.Insert(u))
}

/**
删除用户角色关联关系
*/

func DeleteUserRole(id string) int {
	u := UserRole{Id: id}
	return CheckInt64(config.EngDb.Where("id = ? ", id).Delete(u))
}

/**
修改用户角色关联关系
*/

func UpdateUserRole(uid, rid, remark string) int {
	r := UserRole{Id: commons.EncodeMd5(commons.StringJoin(uid, rid)), Uid: uid, Remark: remark, Rid: rid}
	return CheckInt64(config.EngDb.Where("id = ? ", r.Id).Cols("remark", "uid", "rid").Update(r))

}

/**
查询user的时候带上这个，如果这个里面找不到内容就是null
*/
