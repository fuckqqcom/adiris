package admin

import "time"

type at struct {
	//Id         string //主键id
	//IsDel      int
	//Remark     string
	//Status     int
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}
