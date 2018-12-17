package admin

type Role struct {
	Id     string //主键id
	IsDel  int
	Remark string
	Status int
	Rid    string
	Name   string
	at     `xorm:"extends"`
}
