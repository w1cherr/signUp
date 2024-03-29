package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"signUp/service"
	"github.com/astaxie/beego"
)

func init()  {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:wakeup@tcp(127.0.0.1:3306)/signUp?charset=utf8", 30)
	orm.RegisterModel(
		new(User),
		new(Admin),
		new(Activity),
		new(Swiper),
		new(Message),
		new(SignUp),
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

func CountObjects(qs orm.QuerySeter) (int64, error) {
	cnt, err := qs.Count()
	if err != nil {
		beego.Error("models.CountObjects ", err)
		return 0, err
	}
	return cnt, err
}

func ListObjects(qs orm.QuerySeter, objs interface{}) (int64, error) {
	nums, err := qs.RelatedSel().All(objs)
	if err != nil {
		beego.Error("models.ListObjects ", err)
		return 0, err
	}
	return nums, err
}