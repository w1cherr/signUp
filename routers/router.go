package routers

import (
	"signUp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.CommonController{}, &controllers.WebController{})
}
