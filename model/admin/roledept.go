package admin

type RoleDept struct {
	Rid string
	Did string
	T   `xorm:"extends"`
}
