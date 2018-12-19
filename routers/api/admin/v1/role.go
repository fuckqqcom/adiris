package v1

import (
	"adiris/pkg/app"
	"adiris/service/admin"
	"adiris/tools/common"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

func Role(c app.Eng) {
	var rl admin.RL

	r := app.R{c}
	m := make(map[string]interface{})
	validate = validator.New()

	if !commons.CheckErr(c.ReadJSON(&rl), "bindJson") || !commons.CheckErr(validate.Struct(rl), "validateJson") {
		m["code"] = 500
		r.Response(m)
		return
	}
	switch rl.Flag {
	case "add":
		m["code"] = rl.AddRole()
	case "update":
		m["code"] = rl.UpdateRole()
	case "delete":
		m["code"] = rl.DeleteRole()
	default:
		m["code"] = 500

	}
	fmt.Println(rl)

	r.Response(m)
}

func RoleList(c app.Eng) {
	var p admin.Pn

	r := app.R{c}
	m := make(map[string]interface{})
	validate = validator.New()

	if !commons.CheckErr(c.ReadJSON(&p), "bindJson") || !commons.CheckErr(validate.Struct(p), "validateJson") {
		m["code"] = 500
		r.Response(m)
		return
	}

	m["data"] = p.FindRole()
	r.Response(m)
}
