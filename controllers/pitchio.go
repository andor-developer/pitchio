package controllers

import (
	"github.com/astaxie/beego"
)

type PitchIOController struct {
	beego.Controller
}

func (c *PitchIOController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
