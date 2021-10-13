package main

import (
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/iMikio/SVNChangeLogToCSV-web/routers"
)

func main() {
	beego.Run()
}

func init() {
	runmode := beego.BConfig.RunMode
	if runmode == "prod" {
		config := `{
			"filename":"logs/appname.log",
			"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info"],
			"maxdays":30
			}`
		config = strings.Replace(config, "appname", beego.BConfig.AppName, 1)
		logs.SetLogger(logs.AdapterMultiFile, config)
		logs.SetLevel(logs.LevelInfo)
	}
	logs.Async(1e3)
	logs.Info("runmode is %v", runmode)
}
