package controllers

import (
	"github.com/astaxie/beego"
	"signUp/models"
)

type Common struct {
	beego.Controller
	User models.User
	IsLogin 	bool
}

func (this *Common) ReturnJson(status int, message string, args ...interface{})  {
	result := make(map[string]interface{})
	result["status"] = status
	result["message"] = message

	key := ""

	for _, arg := range args {
		switch arg.(type) {
		case string:
		key = arg.(string)
		default:
		result[key] = arg
		}
	}

	this.Data["json"] = result
	this.ServeJSON()
	this.StopRun()
}

func (this *Common) ReturnSuccess(args ...interface{}) {
	result := make(map[string]interface{})
	result["status"] = 10000
	result["message"] = "success"
	key := ""
	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)
		default:
			result[key] = arg
		}
	}
	this.Data["json"] = result
	this.ServeJSON()
	this.StopRun()
}