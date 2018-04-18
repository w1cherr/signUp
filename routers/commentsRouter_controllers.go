package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["signUp/controllers:WebController"] = append(beego.GlobalControllerRouter["signUp/controllers:WebController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/*`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

}
