package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"signUp/service"
)

func init()  {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:wakeup@tcp(127.0.0.1:3306)/signUp?charset=utf8", 30)
	orm.RegisterModel(
		new(User),
		new(Admin),
		)
	orm.RunSyncdb("default", false, true)
}

func AddAdmin() {
	var admin Admin
	fmt.Println("请输入管理员账号:")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println("请输入管理员密码:")
	var password string
	fmt.Scanf("%s", &password)
	admin.Name = name
	admin.Password = service.StrToMD5(password)
	fmt.Println(admin)
	err := admin.InsertAdmin()
	if err!=nil {
		fmt.Println("发生错误，请重试")
		return
	}
	fmt.Println("创建超级管理员账号成功！")
}