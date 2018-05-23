package admin

import (
	"signUp/models"
	"signUp/controllers"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type AdminSwiperController struct {
	controllers.Common
}

// @router /api/admin/swiper/save [*]
func (this *AdminSwiperController) SaveActivity ()  {
	admin := this.GetSession("admin")
	if admin == nil{
		this.ReturnJson(10403,"请先登录")
		return
	}
	id, _ := this.GetInt64("id")
	cover := this.GetString("cover")
	activityId,_ := this.GetInt64("activityId")
	if len(cover)==0{
		this.ReturnJson(10003, "need all params")
		return
	}
	var swiper models.Swiper
	swiper.Cover = cover
	swiper.ActivityId = activityId
	if id == 0 {
		err := swiper.Insert()
		if (err != nil) {
			beego.Error(err)
			this.ReturnJson(10001, "swiper insert error")
			return
		}
		this.ReturnSuccess()
	} else {
		t := models.Swiper{Id:id}
		t.Read()
		swiper.Id = id
		swiper.CreatedTime = t.CreatedTime
		err := swiper.Update()
		if (err != nil) {
			beego.Error(err)
			this.ReturnJson(10001, "swiper update error")
			return
		}
		this.ReturnSuccess()
	}
}

// @router /api/admin/swiper/list [*]
func (this *AdminSwiperController) ListSwiper() {
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

// @router /api/admin/swiper/delete [*]
func (this *AdminSwiperController) DelateSwiper() {
	id ,_:= this.GetInt64("id")
	swiper := models.Swiper{Id:id}
	if err:=swiper.Read();err!=nil{
		this.ReturnJson(10001,"not found this activity")
		return
	}
	if _,err := swiper.Delete(); err != nil {
		this.ReturnJson(10002,"delete error")
		return
	}
	this.ReturnSuccess()
}

