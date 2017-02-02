// @APIVersion 1.0.0
// @Title MallApi支持的接口
// @Description MallApi提供CMS的API
// @Contact z29759@gmail.com
// @TermsOfServiceUrl http://www.mallapi.com/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/zerodollar/mall-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/SYS_MENU_INF",
			beego.NSInclude(
				&controllers.SYSMENUINFController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
