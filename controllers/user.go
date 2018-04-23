package controllers

import (
	"signUp/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"signUp/service"
	"fmt"
)

type UserController struct {
	Common
}

// @router /api/user/created [*]
func (this *UserController) UserInster()  {
	name := this.GetString("Acount")
	password1 := this.GetString("Password1")
	password2 := this.GetString("Password2")

	//检查是否重复注册
	var errUser models.User
	orm.NewOrm().QueryTable("user").Filter("name", name).One(&errUser)
	if (errUser.Id != 0) {
		this.ReturnJson(10001, "name has resgisted")
		return
	}
	if (len(name)<=0) {
		this.ReturnJson(10002, "name can't empty")
		return
	}
	if (len(password1)<6) {
		this.ReturnJson(10003, " the length of password must max than 6")
		return
	}
	if (len(password1) == 0 || len(password2) == 0) {
		this.ReturnJson(10004, "one password empty")
		return
	}
	var user models.User
	user.Name = name
	user.Password = service.StrToMD5(password1)
	err := user.InsertUser()
	if (err != nil) {
		beego.Error(err)
		this.ReturnJson(10005, "user insert error")
		return
	}

	this.ReturnSuccess()
}

// @router /api/user/login [*]
func (this *UserController) UserLogin()  {
	name := this.GetString("Account")
	password := this.GetString("Password")
	fmt.Println(name)

	user := models.User{}

	err := orm.NewOrm().QueryTable("user").Filter("name", name).One(&user)

	if err != nil {
		this.ReturnJson(10001, "user not found")
		return
	}

	if (user.Password == service.StrToMD5(password)) {
		fmt.Println("当前的user:")
		fmt.Println(user)
		this.SetSession("user", user)
		fmt.Println("当前的session:")
		fmt.Println(this.CruSession)
		this.ReturnSuccess()
	} else {
		this.ReturnJson(10002, "password error")
		return
	}
}

// @router /api/user/logout [*]
func (this *UserController) UserLogout() {
	this.DelSession("user")
	//this.ReturnSuccess()
}