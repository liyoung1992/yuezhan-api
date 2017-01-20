package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["yuezhan-api/controllers:BookController"] = append(beego.GlobalControllerRouter["yuezhan-api/controllers:BookController"],
		beego.ControllerComments{
			"BookParkingLots",
			`/book`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["yuezhan-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yuezhan-api/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["yuezhan-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yuezhan-api/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["yuezhan-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yuezhan-api/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["yuezhan-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yuezhan-api/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["yuezhan-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yuezhan-api/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["yuezhan-api/controllers:UserController"] = append(beego.GlobalControllerRouter["yuezhan-api/controllers:UserController"],
		beego.ControllerComments{
			"Register",
			`/register`,
			[]string{"post"},
			nil})

	//	beego.GlobalControllerRouter["yuezhan-api/controllers:UserController"] = append(beego.GlobalControllerRouter["yuezhan-api/controllers:UserController"],
	//		beego.ControllerComments{
	//			"Login",
	//			`/login`,
	//			[]string{"post"},
	//			nil})

	//	beego.GlobalControllerRouter["yuezhan-api/controllers:UserController"] = append(beego.GlobalControllerRouter["yuezhan-api/controllers:UserController"],
	//		beego.ControllerComments{
	//			"Logout",
	//			`/logout`,
	//			[]string{"get"},
	//			nil})

	//	beego.GlobalControllerRouter["yuezhan-api/controllers:UserController"] = append(beego.GlobalControllerRouter["yuezhan-api/controllers:UserController"],
	//		beego.ControllerComments{
	//			"Hello",
	//			`/:uid/hello`,
	//			[]string{"get"},
	//			nil})

}
