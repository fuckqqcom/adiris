package admin

import (
	"adiris/pkg/config"
	"adiris/pkg/e"
	"adiris/tools/common"
	"adiris/tools/mysql"
)

type User struct {
	Id       string //主键id
	IsDel    int
	Remark   string
	Status   int
	Account  string //账号
	Password string //密码
	Email    string //邮箱
	Phone    string //手机号码
	T        `xorm:"extends"`
}

/**
用户登录
*/

func Login(name, pwd, gid string) {

}

/**
用户注册:
account string   用户名
pwd  string   密码
gid  string   机构id(组id)
*/

func Register(account, pwd, gid string) int {
	uid := commons.EncodeMd5(commons.StringJoin(account, gid))
	pw := commons.EncodeMd5(commons.StringJoin(pwd, account, gid))

	if GetUserByUid(uid) {
		return e.UserExist
	}
	if !GetGidExist(gid) {
		return e.GroupNotExist
	}
	u := User{Account: account, Password: pw, Id: uid, IsDel: 1, Status: 0}
	ug := UserGroup{Uid: uid, Gid: gid, Id: commons.EncodeMd5(commons.StringJoin(uid, gid)), IsDel: 1, Status: 0}
	//给一个默认的角色id
	ur := UserRole{Uid: uid, Rid: "1111", Id: commons.EncodeMd5(commons.StringJoin(uid, "111")), IsDel: 1, Status: 0}

	s := config.EngDb.NewSession()
	var err error
	//这里是不是session多关闭了一次
	defer s.Clone()
	if err = s.Begin(); err != nil {
		return e.Error
	}

	defer s.Rollback()
	if !(mysql.CheckErrInsert(s, u) && mysql.CheckErrInsert(s, ug) && mysql.CheckErrInsert(s, ur)) {
		return e.UserRegisterFail
	}
	s.Commit()
	s.Close()

	return e.Success
}

/**
查找用户是否存在 true：存在 false 不存在
*/
func GetUserByUid(uid string) bool {
	exist, err := config.EngDb.Where("id = ? ", uid).Exist(&User{})

	if commons.CheckErr(err, exist) && exist {
		return true
	}
	return false
}
