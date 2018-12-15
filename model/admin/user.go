package admin

type User struct {
	Account  string //账号
	Password string //密码
	Salt     string //盐
	Email    string //邮箱
	Phone    string //手机号码
	T        `xorm:"extends"`
}
