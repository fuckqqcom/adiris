package admin

type Menu struct {
	Id       string //主键id
	IsDel    int
	Remark   string
	Status   int
	Name     string //菜单名称
	ParentId string //父菜单Id，一级菜单为0
	Url      string //菜单url,类型：1.普通页面（如用户管理， /sys/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀+目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)'
	//Perms    string
	Type int //0:目录(一级tab) 1:菜单 2:按钮
	at   `xorm:"extends"`
}
