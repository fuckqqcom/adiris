package admin

import "adiris/model/admin"

type User struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
	Gid      string `json:"gid" validate:"required" `
}

func (u User) Register() int {
	return admin.Register(u.Account, u.Password, u.Gid)
}
