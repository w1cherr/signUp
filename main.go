package main

import (
	_ "signUp/routers"
	"signUp/models"
	"github.com/astaxie/beego"
	"os"
)
func init() {
	initArgs()
}
func main() {
	beego.Run()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-addAdmin" {
			models.AddAdmin()
		}
	}
}
