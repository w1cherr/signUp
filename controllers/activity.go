package controllers

import (
	"signUp/models"
	"github.com/astaxie/beego/orm"
)

type ActivityController struct {
	Common
}

//@router /api/activity/list [*]
func (this *ActivityController) ListActivity() {
	per, _ := this.GetInt("per")
	page, _ := this.GetInt("page")
	if (per < 1) {
		per = 10
	}
	if(page < 1) {
		page = 1
	}
	qs := orm.NewOrm().QueryTable("activity").Filter("Status",1)
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

//@router /api/activity/detail [*]
func (this *ActivityController) DetailActivity()  {
	id ,_ := this.GetInt64("id")
	activity := models.Activity{}

	err := orm.NewOrm().QueryTable("activity").Filter("id", id).One(&activity)

	if err != nil {
		this.ReturnJson(10001, "activity not found")
		return
	}
	res := make(map[string]interface{})
	res["status"] = 10000
	res["data"] = activity
	this.Data["json"] = res
	this.ServeJSON()
	return
}