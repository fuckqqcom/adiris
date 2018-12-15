package admin

type Log struct {
	UserName string
	Uri      string
	Method   string
	Ip       string
	T        `xorm:"extends"`
}
