package controllers

import (
	"bytes"
	"strconv"

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
	// limit for 1M
	if TextBytes := len(s); TextBytes > 1048576 {
		logs.Warning("I don't support to sizes larger than 1MB text. size:%v", TextBytes)
		c.Data["json"] = "I don't support to sizes larger than 1MB text."
		c.Abort("400")
	}
	parsed, err := models.ParseChangeLog(&s)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = err.Error()
		c.Abort("500")
	}

	rw := c.Ctx.ResponseWriter
	rw.Header().Set("Content-Type", "text/tab-separated-values")
	rw.Header().Set("Content-Disposition", "attachment; filename=changelog.tsv")
	// rw.Header().Set("Content-Type", "text/csv")
	// rw.Header().Set("Content-Disposition", "attachment; filename=changelog.csv")

	bf := &bytes.Buffer{}
	models.WriteCsv(bf, parsed)
	rw.Header().Set("Content-Length", strconv.Itoa(bf.Len()))
	bf.WriteTo(rw)
}
