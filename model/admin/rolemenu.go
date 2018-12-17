package admin

type RoleMenu struct {
	Id     string //主键id
	IsDel  int
	Remark string
	Status int
	Rid    string
	Mid    string
	at     `xorm:"extends"`
}
