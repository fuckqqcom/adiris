package admin

type RoleGroup struct {
	Rid string
	Did string
	T   `xorm:"extends"`
}
