package v1

import (
	"adiris/pkg/app"
	"adiris/service/admin"
	"adiris/tools/common"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func Register(c app.Eng) {
	var u admin.User
	validate = validator.New()
	r := app.R{c}
	m := make(map[string]interface{})
	if !commons.CheckErr(c.ReadJSON(&u), "bindJson") || !commons.CheckErr(validate.Struct(u), "validateJson") {
		m["code"] = 500
		r.Response(m)
		return
	}
	m["code"] = u.Register()
	r.Response(m)
}

func GetUserPer(c app.Eng) {
	r := app.R{c}
	m := make(map[string]interface{})
	var u admin.User

	u.GetUserPer()
	r.Response(m)
}
