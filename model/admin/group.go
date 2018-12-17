package admin

import (
	"adiris/pkg/config"
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
