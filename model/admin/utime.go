package admin

import (
	"adiris/pkg/e"
	"adiris/tools/common"
	"time"
)

type at struct {
	//Id         string //主键id
	//IsDel      int
	//Remark     string
	//Status     int
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

func CheckBool(exist bool, err error) bool {
	if commons.CheckErr(err, exist) && exist {
		return true
	} else {
		return false
	}
}

func CheckInt64(exist int64, err error) int {
	if commons.CheckErr(err, exist) && exist != 0 {
		return e.Success
	} else {
		return e.Error
	}
}
