package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["signUp/controllers/admin:ActivityController"] = append(beego.GlobalControllerRouter["signUp/controllers/admin:ActivityController"],
		beego.ControllerComments{
			Method: "ListActivity",
			Router: `/api/admin/activity/list`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["signUp/controllers/admin:ActivityController"] = append(beego.GlobalControllerRouter["signUp/controllers/admin:ActivityController"],
		beego.ControllerComments{
			Method: "SaveActivity",
			Router: `/api/admin/activity/save`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["signUp/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["signUp/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "AdminLogin",
			Router: `/api/admin/login`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["signUp/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["signUp/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "AdminLogout",
			Router: `/api/admin/logout`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

}
