package controllers

import (
	//"github.com/astaxie/beego"
)

type WebController struct {
	CommonController
}

// @router /* [*]
func (this *WebController) Index() {
	this.TplName = "mobile/index.html"
}