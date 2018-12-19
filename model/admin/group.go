package admin

import (
	"adiris/pkg/config"
	"adiris/pkg/e"
	"adiris/tools/common"
)

type Group struct {
	Id       string //主键id
	IsDel    int
	Remark   string
	Status   int
	Name     string //机构名称
	ParentId string //上级机构Id ,一级机构为0
	OrderNum int    //排序
	at       `xorm:"extends"`
}

/**
添加机构(组)
*/

func AddGroup(name, remark string, status int) int {
	g := Group{Id: commons.EncodeMd5(name), IsDel: 1, Name: name, Status: status, Remark: remark}

	if GetGidExistTb(g.Id) {
		return e.RoleExist
	}
	return CheckInt64(config.EngDb.Insert(g))
}

/**
删除机构
*/

func DeleteGroup(id string) int {
	g := Group{Id: id}
	return CheckInt64(config.EngDb.Where("id = ?", id).Delete(g))
}

/**
修改机构
*/

func UpdateGroup(name, remark string, status int) int {
	r := Role{Id: commons.EncodeMd5(name), IsDel: 1, Name: name, Status: status, Remark: remark}
	return CheckInt64(config.EngDb.Where("id = ? ", r.Id).Cols("remark", "status", "name").Update(r))
}

/**
查询机构
*/

func GetGroupList(pn, ps int) interface{} {
	var g []Group
	count, err := config.EngDb.Where("is_del = 1 and status = 1").Desc("create_time").Limit(ps, (pn-1)*ps).FindAndCount(&g)
	CheckInt64(count, err)
	m := make(map[string]interface{})
	m["count"] = int(count)
	m["data"] = g
	return m
}

/**
检查组织结构是否存在
*/
func GetGidExist(id string, c chan map[string]interface{}) {
	exist, err := config.EngDb.Where("id = ? and Parent_id = '0' ", id).Exist(&Group{})
	m := make(map[string]interface{})
	if commons.CheckErr(err, exist) && exist {
		m["flag"] = true
		c <- m
	} else {
		m["flag"] = false
		c <- m
	}
}

func GetGidExistTb(id string) bool {
	exist, err := config.EngDb.Where("id = ? and Parent_id = '0' ", id).Exist(&Group{})
	//m := make(map[string]interface{})
	if commons.CheckErr(err, exist) && exist {
		return true
	} else {
		return false
	}
}
