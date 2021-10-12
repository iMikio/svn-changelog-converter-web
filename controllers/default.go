package controllers

import (
	"encoding/csv"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/iMikio/SVNChangeLogToCSV-web/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	s := c.GetString("text")
	logs.Debug(s)
	parsed, err := models.ParseChangeLog(&s)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = err.Error()
		c.Abort("500")
	}
	logs.Debug(*parsed)

	rw := c.Ctx.ResponseWriter
	rw.Header().Set("Content-Type", "text/csv")
	rw.Header().Set("Content-Disposition", "attachment; filename=changelog.csv")
	cw := csv.NewWriter(rw)
	defer cw.Flush()

	cw.WriteAll(*parsed)
}
