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
			"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],
			"maxdays":90
			}`
		config = strings.Replace(config, "appname", beego.BConfig.AppName, 1)
		logs.SetLogger(logs.AdapterMultiFile, config)
	}
	logs.Async()
	logs.Info("runmode is %v", runmode)
}
