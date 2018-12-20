package admin

import (
	"adiris/model/admin"
	"adiris/tools/common"
	"fmt"
	"time"
)

type User struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
	Gid      string `json:"gid" validate:"required" `
}

func (u User) Register() int {
	//模拟实战 ，这里随机

	a := commons.RandStringBytes(4) + fmt.Sprintf("%s", time.Now().UnixNano())
	return admin.Register(u.Account+a, u.Password+a, u.Gid)
}

func (u User) GetUserPer() {
	admin.GetUserPer()
}
