package controllers

import (
	//"github.com/astaxie/beego"
)

type WebController struct {
	Common
}

// @router /* [*]
func (this *WebController) Index() {
	this.TplName = "mobile/index.html"
}

// @router /admin/* [*]
func (this *WebController) Admin() {
	this.TplName = "admin/index.html"
}