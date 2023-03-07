package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	IsLogin   bool
	LoginUser interface{}
	beego.Controller
}

func (this *BaseController) Prepare() {
	user := this.GetSession("loginuser")
	fmt.Println("loginuser------>", user)
	if user != nil {
		this.IsLogin = true
		this.LoginUser = user
	} else {
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}
