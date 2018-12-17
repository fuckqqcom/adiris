package admin

type UserRole struct {
	Id     string `xorm:"varchar(100) notnull pk index unique 'id'"` //主键id
	IsDel  int
	Remark string
	Status int
	Uid    string
	Rid    string
	at     `xorm:"extends"`
}
