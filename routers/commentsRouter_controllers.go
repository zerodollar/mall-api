package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/zerodollar/mall-api/controllers:SYSMENUINFController"] = append(beego.GlobalControllerRouter["github.com/zerodollar/mall-api/controllers:SYSMENUINFController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/zerodollar/mall-api/controllers:SYSMENUINFController"] = append(beego.GlobalControllerRouter["github.com/zerodollar/mall-api/controllers:SYSMENUINFController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/zerodollar/mall-api/controllers:SYSMENUINFController"] = append(beego.GlobalControllerRouter["github.com/zerodollar/mall-api/controllers:SYSMENUINFController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/zerodollar/mall-api/controllers:SYSMENUINFController"] = append(beego.GlobalControllerRouter["github.com/zerodollar/mall-api/controllers:SYSMENUINFController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/zerodollar/mall-api/controllers:SYSMENUINFController"] = append(beego.GlobalControllerRouter["github.com/zerodollar/mall-api/controllers:SYSMENUINFController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
