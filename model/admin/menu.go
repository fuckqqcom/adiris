package admin

import (
	"adiris/pkg/config"
	"adiris/pkg/e"
	"adiris/tools/common"
)

type Menu struct {
	Id       string //主键id
	IsDel    int
	Remark   string
	Status   int
	Name     string //菜单名称
	ParentId string //父菜单Id，一级菜单为0
	Url      string //菜单url,类型：1.普通页面（如用户管理， /sys/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀+目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)'
	//Perms    string
	Type int //0:目录(一级tab) 1:菜单 2:按钮
	at   `xorm:"extends"`
}

/**
添加菜单/todo
*/

func AddMenu(name, remark string, status int) int {
	m := Menu{Id: commons.EncodeMd5(name), IsDel: 1, Name: name, Status: status, Remark: remark}

	//if GetGidExistTb(m.Id) {
	//	return e.RoleExist
	//}
	return CheckInt64(config.EngDb.Insert(m))
}

/**
删除机构
*/

func DeleteMenu(id string) int {
	m := Menu{Id: id}
	return CheckInt64(config.EngDb.Where("id = ?", id).Delete(m))
}

/**
修改机构
*/

func UpdateMenu(name, remark string, status int) int {
	m := Menu{Id: commons.EncodeMd5(name), IsDel: 1, Name: name, Status: status, Remark: remark}
	return CheckInt64(config.EngDb.Where("id = ? ", r.Id).Cols("remark", "status", "name").Update(m))
}

/**
查询机构
*/

func GetMenuList(pn, ps int) interface{} {
	var g []Menu
	count, err := config.EngDb.Where("is_del = 1 and status = 1").Desc("create_time").Limit(ps, (pn-1)*ps).FindAndCount(&g)
	CheckInt64(count, err)
	m := make(map[string]interface{})
	m["count"] = int(count)
	m["data"] = g
	return m
}
