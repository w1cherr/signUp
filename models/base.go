package models

import (
	"github.com/astaxie/beego/orm"
)

func init()  {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:wakeup@tcp(127.0.0.1:3306)/signUp?charset=utf8", 30)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}