package admin

import (
	"signUp/models"
	"signUp/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AdminUserController struct {
	controllers.Common
}

// @router /api/admin/user/list [*]
func (this *AdminUserController) ListUser() {
	admin := this.GetSession("admin")
	if admin == nil{
		this.ReturnJson(10403,"请先登录")
		return
	}
	per, _ := this.GetInt("per")
	page, _ := this.GetInt("page")
	if (per < 1) {
		per = 10
	}
	if(page < 1) {
		page = 1
	}
	qs := orm.NewOrm().QueryTable("user")
	cnt, _ := models.CountObjects(qs)
	var total int
	if (int(cnt) % per > 0) {
		total = (int(cnt)/per) + 1
	} else {
		total = (int(cnt)/per)
	}
	hasNext := false
	if (page < total) {
		hasNext = true
	}
	qs = qs.OrderBy("-created_time").Limit(per, (page-1)*per).RelatedSel()
	var user []models.User
	models.ListObjects(qs, &user)
	this.ReturnSuccess("data",user,"page",page,"hasNext",hasNext,"cnt",cnt,"per",per,"total", total)
}

// @router /api/admin/user/handle [*]
func (this *AdminUserController) HandleUser() {
	id ,_:= this.GetInt64("id")
	status,_ := this.GetInt("status")
	user := models.User{Id:id}
	if err:=user.Read();err!=nil{
		this.ReturnJson(10001,"not found this user")
		return
	}
	user.CanSignUp = status
	err := user.Update()
	if (err != nil) {
		beego.Error(err)
		this.ReturnJson(10001, "user update error")
		return
	}
	this.ReturnSuccess()
}

