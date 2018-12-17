package admin

type UserGroup struct {
	Id     string //主键id
	IsDel  int
	Remark string
	Status int
	Uid    string
	Gid    string
	at     `xorm:"extends"`
}
