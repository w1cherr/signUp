package routers

import (
	"signUp/controllers"
	"signUp/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(
		&controllers.Common{},
		&controllers.WebController{},
		&controllers.UserController{},
		&controllers.SwiperController{},
		&controllers.ActivityController{},
		&admin.AdminController{},
		&admin.ActivityController{},
		&admin.AdminUserController{},
		&admin.AdminSwiperController{},
		)
	beego.SetStaticPath("/upload/file", "upload/file")
}
