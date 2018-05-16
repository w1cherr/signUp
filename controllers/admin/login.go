package admin

import (
	"signUp/models"
	"signUp/service"
	"fmt"
	"github.com/astaxie/beego/orm"
)

// @router /api/admin/login [*]
func (this *AdminController) AdminLogin()  {
	name := this.GetString("Username")
	password := this.GetString("Password")
	fmt.Println(name)

	admin := models.Admin{}

	err := orm.NewOrm().QueryTable("admin").Filter("name", name).One(&admin)

	if err != nil {
		this.ReturnJson(10001, "admin not found")
		return
	}

	if (admin.Password == service.StrToMD5(password)) {
		fmt.Println("当前的admin:")
		fmt.Println(admin)
		this.SetSession("admin", admin)
		fmt.Println("当前的session:")
		fmt.Println(this.CruSession)
		this.ReturnSuccess()
	} else {
		this.ReturnJson(10002, "password error")
		return
	}
}

// @router /api/admin/logout [*]
func (this *AdminController) AdminLogout() {
	this.DelSession("admin")
	this.ReturnSuccess()
}