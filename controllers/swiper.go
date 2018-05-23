package controllers

import (
	"signUp/models"
	"github.com/astaxie/beego/orm"
)

type SwiperController struct {
	Common
}

// @router /api/swiper/list [*]
func (this *SwiperController) ListSwiper() {
	per, _ := this.GetInt("per")
	page, _ := this.GetInt("page")
	if (per < 1) {
		per = 10
	}
	if(page < 1) {
		page = 1
	}
	qs := orm.NewOrm().QueryTable("swiper")
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
	var swiper []models.Swiper
	models.ListObjects(qs, &swiper)
	this.ReturnSuccess("data",swiper,"page",page,"hasNext",hasNext,"cnt",cnt,"per",per,"total", total)
}
