package admin

type Log struct {
	Id       string //主键id
	IsDel    int
	Remark   string
	Status   int
	UserName string
	Uri      string
	Method   string
	Ip       string
	at       `xorm:"extends"`
}
