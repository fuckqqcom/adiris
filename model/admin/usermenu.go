package admin

type UserMenu struct {
	Id     string //主键id
	IsDel  int
	Remark string
	Status int
	Uid    string
	Mid    string
	at     `xorm:"extends"`
}
