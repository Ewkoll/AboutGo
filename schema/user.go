package schema

import (
	"time"
)

type User struct {
	ID          string    `orm:"column(id);pk;description(用户编号)" json:""`
	RoleID      string    `orm:"column(role_id);description(用户角色信息)"`
	UserName    string    `orm:"column(user_name);unique;description(用户名称)"`
	Password    string    `orm:"column(password);description(用户密码)"`
	RealName    string    `orm:"column(real_name);description(昵称)"`
	IdCard      string    `orm:"column(id_card);description(身份证)"`
	Phone       string    `orm:"column(phone);description(电话号码)"`
	Email       string    `orm:"column(email);description(邮箱地址)"`
	Salt        string    `orm:"column(salt);description(密码Salt)"`
	LastLogin   string    `orm:"column(last_login);null;description(最后登录)"`
	LastIp      string    `orm:"column(last_ip);null;description(最后登录IP)"`
	Status      uint32    `orm:"column(status);description(状态)"`
	CreatedTime time.Time `orm:"auto_now_add;type(datetime);description(创建时间)"`
	UpdatedTime time.Time `orm:"auto_now;type(datetime);description(修改时间)"`
}
