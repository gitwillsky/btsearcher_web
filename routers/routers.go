package routers
import (
	"github.com/astaxie/beego"
	"github.com/gitwillsky/btsearcher.web/controllers/index"
)


func init() {
	beego.Include(
		// front index
		&index.IndexController{},
	)
}