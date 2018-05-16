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
		&admin.AdminController{},
		)
}
