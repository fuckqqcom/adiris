package admin

type RoleMenu struct {
	Rid string
	Mid string
	T   `xorm:"extends"`
}
