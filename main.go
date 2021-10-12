package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/iMikio/SVNChangeLogToCSV-web/routers"
)

func main() {
	beego.SetStaticPath("/bootstrap", "static/lib/bootstrap-5.0.2-dist")
	beego.Run()
}
