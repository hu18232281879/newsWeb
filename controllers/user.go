package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"newsWeb/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) ShowRegister() {
	this.TplName = "register.html"
}
func (this *UserController) HandleRegister() {
	userName := this.GetString("userName")
	password := this.GetString("password")
	if userName == "" || password == "" {
		this.Data["err"] = "用户名或密码不能为空"
		this.TplName = "register.html"
		return
	}
	o := orm.NewOrm()
	user := new(models.User)
	user.Name = userName
	user.Pwd = password
	o.Insert(user)
	this.Redirect("/login", 302)
}
func (this *UserController) ShowLogin() {
	this.TplName = "login.html"
}
func (this *UserController) HandleLogin() {
	userName := this.GetString("userName")
	password := this.GetString("password")
	if userName == "" || password == "" {
		this.Data["errmsg"] = "用户名,密码不能为空!"
		this.TplName = "login.html"
		return
	}
	o := orm.NewOrm()
	user := new(models.User)
	user.Name = userName
	err := o.Read(user, "Name")
	if err != nil {
		this.Data["errmsg"] = "用户名或密码错误"
		this.TplName = "login.html"
		return
	}
	user.Pwd = password
	err = o.Read(user, "Pwd")
	if err != nil {
		this.Data["errmsg"] = "用户名或密码错误"
		this.TplName = "login.html"
		return
	}
	this.Redirect("/index", 302)
}

