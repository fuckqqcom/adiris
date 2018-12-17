package admin

type UserRole struct {
	Id     string //主键id
	IsDel  int
	Remark string
	Status int
	Uid    string
	Rid    string
	at     `xorm:"extends"`
}
