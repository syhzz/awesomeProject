package controllers

import (
	"BeegoDemo/models"
	"BeegoDemo/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:"+username, "password:"+password)
	id := models.QueryUserWithParam(username, utils.MD5(password))
	fmt.Println("id:", id)
	if id > 0 {
		this.SetSession("loginuser", username)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "登入成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登入失败"}
	}
	this.ServeJSON()
}
