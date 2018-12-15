package admin

type UserRole struct {
	Uid string
	Rid string
	T   `xorm:"extends"`
}
