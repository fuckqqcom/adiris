package admin

type User struct {
	Uid        string //用户id
	Account    string //账号
	Password   string //密码
	Salt       string //盐
	Email      string //邮箱
	Phone      string //手机号码
	Status     int    //'状态  -1：禁用   1：正常'
	DeptId     string //机构id
	Remark     string //标注
	CreateTime string //创建时间
	UpdateTime string //修改时间
	IsDel      int    //是否删除
}
