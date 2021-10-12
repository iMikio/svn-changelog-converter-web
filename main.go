package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/iMikio/SVNChangeLogToCSV-web/routers"
)

func main() {
	beego.Run()
}
