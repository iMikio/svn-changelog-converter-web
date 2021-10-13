package routers

import (
	"github.com/iMikio/svn-changelog-converter-web/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
