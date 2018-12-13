package admin

import "time"

type T struct {
	Id         string //主键id
	IsDel      int
	Remark     string
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}
