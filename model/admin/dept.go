package admin

type Dept struct {
	Name     string //机构名称
	ParentId string //上级机构Id ,一级机构为0
	OrderNum int    //排序
	T
}
