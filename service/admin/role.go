package admin

import "adiris/model/admin"

type RL struct {
	Id     string `json:"id" validate:"required"`
	Flag   string `json:"flag" validate:"required"` //请求方式
	Remark string `json:"remark" validate:"required"`
	Status int    `json:"status" validate:"required"`
	Name   string `json:"name" validate:"required"`
}

type Pn struct {
	Pn int `json:"pn" validate:"gte=0,lte=10"`
	Ps int `json:"ps" validate:"required"`
}

type Role struct {
	admin.Role
}

func (r RL) AddRole() int {
	return admin.AddRole(r.Name, r.Remark, r.Status)
}

func (r RL) UpdateRole() int {
	return admin.UpdateRole(r.Name, r.Remark, r.Status)
}

func (r RL) DeleteRole() int {
	return admin.DeleteRole(r.Id)
}

func (p Pn) FindRole() interface{} {
	return admin.GetRoleList(p.Pn, p.Ps)
}
