package admin

import (
	"signUp/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ActivityController struct {
	AdminController
}

//@router /api/admin/activity/save [*]
func (this *ActivityController) SaveActivity ()  {
	admin := this.GetSession("admin")
	if admin == nil{
		this.ReturnJson(10403,"请先登录")
		return
	}
	id, _ := this.GetInt64("id")
	title := this.GetString("title")
	introduction := this.GetString("introduction")
	userLimit,_ := this.GetInt64("userLimit")
	status,_ := this.GetInt("status")
	startTime := this.GetString("startTime")
	endTime := this.GetString("endTime")
	imgUrl := this.GetString("imgUrl")
	cover := this.GetString("cover")

	if len(title)==0 || len(introduction) == 0{
		this.ReturnJson(10003, "need all params")
		return
	}
	var activity models.Activity
	activity.Title = title
	activity.Introduction = introduction
	activity.UserLimit = userLimit
	activity.Status = status
	activity.StartTime = startTime
	activity.EndTime = endTime
	activity.ImgUrl = imgUrl
	activity.Cover = cover
	if id == 0 {
		err := activity.Insert()
		if (err != nil) {
			beego.Error(err)
			this.ReturnJson(10001, "activity insert error")
			return
		}
		this.ReturnSuccess()
	} else {
		t := models.Activity{Id:id}
		t.Read()
		activity.Id = id
		activity.CreatedTime = t.CreatedTime
		err := activity.Update()
		if (err != nil) {
			beego.Error(err)
			this.ReturnJson(10001, "activity update error")
			return
		}
		this.ReturnSuccess()
	}
}

//@router /api/admin/activity/list [*]
func (this *ActivityController) ListActivity() {
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
	qs := orm.NewOrm().QueryTable("activity")
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
	var activity []models.Activity
	models.ListObjects(qs, &activity)
	this.ReturnSuccess("data",activity,"page",page,"hasNext",hasNext,"cnt",cnt,"per",per,"total", total)
}

//@router /api/admin/activity/delete [*]
func (this *ActivityController) DelateActivity() {
	id ,_:= this.GetInt64("id")
	activity := models.Activity{Id:id}
	if err:=activity.Read();err!=nil{
		this.ReturnJson(10001,"not found this activity")
		return
	}
	if _,err := activity.Delete(); err != nil {
		this.ReturnJson(10002,"delete error")
		return
	}
	this.ReturnSuccess()
}