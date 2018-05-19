package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["signUp/controllers:ActivityController"] = append(beego.GlobalControllerRouter["signUp/controllers:ActivityController"],
		beego.ControllerComments{
			Method: "DetailActivity",
			Router: `/api/activity/detail`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["signUp/controllers:ActivityController"] = append(beego.GlobalControllerRouter["signUp/controllers:ActivityController"],
		beego.ControllerComments{
			Method: "ListActivity",
			Router: `/api/activity/list`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["signUp/controllers:SwiperController"] = append(beego.GlobalControllerRouter["signUp/controllers:SwiperController"],
		beego.ControllerComments{
			Method: "ListSwiper",
			Router: `/api/swiper/list`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["signUp/controllers:UserController"] = append(beego.GlobalControllerRouter["signUp/controllers:UserController"],
		beego.ControllerComments{
			Method: "UserInster",
			Router: `/api/user/created`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["signUp/controllers:UserController"] = append(beego.GlobalControllerRouter["signUp/controllers:UserController"],
		beego.ControllerComments{
			Method: "UserIsLogin",
			Router: `/api/user/isLogin`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["signUp/controllers:UserController"] = append(beego.GlobalControllerRouter["signUp/controllers:UserController"],
		beego.ControllerComments{
			Method: "UserLogin",
			Router: `/api/user/login`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["signUp/controllers:UserController"] = append(beego.GlobalControllerRouter["signUp/controllers:UserController"],
		beego.ControllerComments{
			Method: "UserLogout",
			Router: `/api/user/logout`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["signUp/controllers:WebController"] = append(beego.GlobalControllerRouter["signUp/controllers:WebController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/*`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["signUp/controllers:WebController"] = append(beego.GlobalControllerRouter["signUp/controllers:WebController"],
		beego.ControllerComments{
			Method: "Admin",
			Router: `/admin/*`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["signUp/controllers:WebController"] = append(beego.GlobalControllerRouter["signUp/controllers:WebController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/api/admin/file/upload`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

}
