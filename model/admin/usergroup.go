package admin

type UserGroup struct {
	Id     string `xorm:"varchar(100) notnull pk index unique 'id'"` //主键id
	IsDel  int
	Remark string
	Status int
	Uid    string
	Gid    string
	at     `xorm:"extends"`
}
