package controllers

import (
	"bytes"
	"strconv"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/iMikio/svn-changelog-converter-web/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	s := c.GetString("text")
	typ := c.GetString("type")
	code := c.GetString("code")
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

	bf := &bytes.Buffer{}
	var comma rune
	if typ == "tsv" {
		comma = rune('\t')
		rw.Header().Set("Content-Type", "text/tab-separated-values")
		rw.Header().Set("Content-Disposition", "attachment; filename=changelog.tsv")
	} else {
		// csv
		comma = rune(',')
		rw.Header().Set("Content-Type", "text/csv")
		rw.Header().Set("Content-Disposition", "attachment; filename=changelog.csv")
	}

	if code == "shift-jis" {
		tw := transform.NewWriter(bf, japanese.ShiftJIS.NewEncoder())
		defer tw.Close()
		models.WriteCsv(tw, parsed, comma, true)
	} else {
		models.WriteCsv(bf, parsed, comma, false)
	}
	rw.Header().Set("Content-Length", strconv.Itoa(bf.Len()))
	bf.WriteTo(rw)
}
