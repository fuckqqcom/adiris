package admin

import (
	"adiris/pkg/config"
	"adiris/tools/common"
)

type UserMenu struct {
	Id     string //主键id
	IsDel  int
	Remark string
	Status int
	Uid    string
	Mid    string
	at     `xorm:"extends"`
}

/**
给用户分配对应的角色
一个用户只能对应一个角色
*/

func AddUserMenu(uid, mid, remark string) int {
	u := UserMenu{Id: commons.EncodeMd5(commons.StringJoin(uid, mid)), IsDel: 1, Remark: remark, Status: 1, Uid: uid, Mid: mid}
	return CheckInt64(config.EngDb.Insert(u))
}

/**
删除用户角色关联关系
*/

func DeleteUserMenu(id string) int {
	u := UserMenu{Id: id}
	return CheckInt64(config.EngDb.Where("id = ? ", id).Delete(u))
}

/**
修改用户角色关联关系
*/

func UpdateUserMeun(uid, mid, remark string) int {
	r := UserMenu{Id: commons.EncodeMd5(commons.StringJoin(uid, mid)), Uid: uid, Remark: remark, Mid: mid}
	return CheckInt64(config.EngDb.Where("id = ? ", r.Id).Cols("remark", "uid", "mid").Update(r))

}
