package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/gitwillsky/btsearcher.web/controllers/index:IndexController"] = append(beego.GlobalControllerRouter["github.com/gitwillsky/btsearcher.web/controllers/index:IndexController"],
		beego.ControllerComments{
			"Index",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gitwillsky/btsearcher.web/controllers/index:IndexController"] = append(beego.GlobalControllerRouter["github.com/gitwillsky/btsearcher.web/controllers/index:IndexController"],
		beego.ControllerComments{
			"List",
			`/list.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/gitwillsky/btsearcher.web/controllers/index:IndexController"] = append(beego.GlobalControllerRouter["github.com/gitwillsky/btsearcher.web/controllers/index:IndexController"],
		beego.ControllerComments{
			"View",
			`/view/:id:int.html`,
			[]string{"get"},
			nil})

}
