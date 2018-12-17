package admin

type RoleGroup struct {
	Id     string //主键id
	IsDel  int
	Remark string
	Status int
	Rid    string
	Did    string
	at     `xorm:"extends"`
}
