package admin

type UserMenu struct {
	Uid string
	Mid string
	T   `xorm:"extends"`
}
