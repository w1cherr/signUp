package controllers

import (
	"signUp/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"signUp/service"
	"fmt"
	"time"
)

type UserController struct {
	Common
}

//@router /api/user/created [*]
func (this *UserController) UserInster() {
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
	if (len(name) <= 0) {
		this.ReturnJson(10002, "name can't empty")
		return
	}
	if (len(password1) < 6) {
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
	user.CanSignUp = 1
	err := user.InsertUser()
	if (err != nil) {
		beego.Error(err)
		this.ReturnJson(10005, "user insert error")
		return
	}

	this.ReturnSuccess()
}

//@router /api/user/login [*]
func (this *UserController) UserLogin() {
	name := this.GetString("Account")
	password := this.GetString("Password")

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
	this.ReturnSuccess()
}

// @router /api/user/isLogin [*]
func (this *UserController) UserIsLogin() {
	userinfo:=this.GetSession("user")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
		return
	}

	user := userinfo.(models.User)
	res := make(map[string]interface{})
	res["status"] = 10000
	res["data"] = user
	this.Data["json"] = res
	this.ServeJSON()
	return
}


// @router /api/user/signup [*]
func (this *UserController) Signup() {

	id, _ := this.GetInt64("id")

	userinfo:=this.GetSession("user")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
		return
	}

	user := userinfo.(models.User)

	if user.CanSignUp != 1 {
		this.ReturnJson(10001, "用户不可报名")
		return
	}

	activity := models.Activity{Id: id}




	var errSignUp models.SignUp
	orm.NewOrm().QueryTable("sign_up").Filter("activity_id",activity.Id).Filter("user_id",user.Id).One(&errSignUp)
	if (errSignUp.Id != 0) {
		this.ReturnJson(10007, "不能重复报名")
		return
	}



	if err := activity.Read(); err != nil {
		this.ReturnJson(10001, "not found this activity")
		return
	}
	startTime, err := time.Parse("2006-01-02 15:04:05", activity.StartTime)
	if err != nil {
		this.ReturnJson(10002, "date parse error")
		return;
	}

	//endTime, err := time.Parse("2006-01-02 15:04:05", activity.EndTime)
	if err != nil {
		this.ReturnJson(10002, "date parse error")
		return;
	}
	startUnix := startTime.Unix()
	//endUnix := endTime.Unix()
	nowUnix := time.Now().Unix()
	if startUnix < nowUnix {
		this.ReturnJson(10003, "活动报名已结束")
		return
	}

	if activity.Status != 1 {
		this.ReturnJson(10004, "活动失效")
		return
	}

	qs := orm.NewOrm().QueryTable("sign_up").Filter("activity_id", id).Filter("status__in", 0, 1)
	userCount, _ := models.CountObjects(qs)
	fmt.Println(activity.UserCount)
	fmt.Println(userCount)
	if activity.UserLimit <= userCount {
		this.ReturnJson(10005, "人数已满")
		return
	}

	o := orm.NewOrm()
	o.Begin()
	var signUp models.SignUp
	signUp.User = &user
	signUp.Activity = &activity
	if _, err := o.Insert(&signUp); err != nil {
		o.Rollback()
		this.ReturnJson(10006, "报名失败")
		return
	}
	activity.UserCount += 1
	if _, err := o.Update(&activity); err != nil {
		o.Rollback()
		this.ReturnJson(10006, "报名失败")
		return
	}
	if err := o.Commit(); err != nil {
		this.ReturnJson(10006, "报名失败")
		return
	}
	this.ReturnSuccess("报名成功")
}

// @router /api/user/activity/list [*]
func (this *UserController) ListUserActivity() {
	userinfo:=this.GetSession("user")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
		return
	}
	user := userinfo.(models.User)
	per, _ := this.GetInt("per")
	page, _ := this.GetInt("page")
	if (per < 1) {
		per = 10
	}
	if(page < 1) {
		page = 1
	}
	qs := orm.NewOrm().QueryTable("sign_up").Filter("user_id", user.Id)
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
	var signUp []models.SignUp
	models.ListObjects(qs, &signUp)
	this.ReturnSuccess("data",signUp,"page",page,"hasNext",hasNext,"cnt",cnt,"per",per,"total", total)
}

// @router /api/user/message/list [*]
func (this *UserController) ListMessage() {
	userinfo:=this.GetSession("user")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
		return
	}
	user := userinfo.(models.User)
	per, _ := this.GetInt("per")
	page, _ := this.GetInt("page")
	if (per < 1) {
		per = 10
	}
	if(page < 1) {
		page = 1
	}
	qs := orm.NewOrm().QueryTable("message").Filter("user_id", user.Id)
	fmt.Println(user.Id)
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
	var message []models.Message
	models.ListObjects(qs, &message)
	this.ReturnSuccess("data",message,"page",page,"hasNext",hasNext,"cnt",cnt,"per",per,"total", total)
}
