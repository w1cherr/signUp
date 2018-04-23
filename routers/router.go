package routers

import (
	"signUp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.Common{}, &controllers.WebController{}, &controllers.UserController{})
}
