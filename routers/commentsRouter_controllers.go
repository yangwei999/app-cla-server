package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:AuthController"],
		beego.ControllerComments{
			Method:           "Auth",
			Router:           "/:platform/:purpose",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:AuthController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/authcodeurl/:platform/:purpose",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAOrgController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAOrgController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAOrgController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAOrgController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/:platform/:org_id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAOrgController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAOrgController"],
		beego.ControllerComments{
			Method:           "GetSigningPageInfo",
			Router:           "/:platform/:org_id/:apply_to",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAOrgController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAOrgController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:uid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAOrgController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:CLAOrgController"],
		beego.ControllerComments{
			Method:           "GetBlankPdf",
			Router:           "/blank-pdf/:cla_org_id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:EmailController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:EmailController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:EmailController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:EmailController"],
		beego.ControllerComments{
			Method:           "Auth",
			Router:           "/auth/:platform",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:EmailController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:EmailController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/authcodeurl/:platform",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:IndividualSigningController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:IndividualSigningController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/:cla_org_id",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:IndividualSigningController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-server/controllers:IndividualSigningController"],
		beego.ControllerComments{
			Method:           "Check",
			Router:           "/:platform/:org/:repo",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
}