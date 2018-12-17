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
	T        `xorm:"extends"`
}

/**
检查组织结构是否存在
*/
func GetGidExist(id string) bool {
	exist, err := config.EngDb.Where("id = ? and Parent_id = 0 ", id).Exist(&Group{})

	if commons.CheckErr(err, exist) && exist {
		return true
	}
	return false
}
