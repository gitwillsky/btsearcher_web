package index
import (
	"github.com/astaxie/beego"
	"github.com/gitwillsky/btsearcher.web/models"
)

type BaseController struct {
	beego.Controller
}

type JSONData struct {
	Code int
	Data interface{}
}

func (this *BaseController) Prepare() {
	// 站点是否关闭
	if ok, _ := beego.AppConfig.Bool("Status"); !ok {
		this.Ctx.WriteString("站点正在升级中,请稍后访问.")
		this.StopRun()
		return
	}

	// xsrf
	//this.XsrfToken()
	//this.Data["xsrfdata"] = template.HTML(this.XsrfFormHtml())
	allCount, _ := models.Count()
	filteredCount, _ := models.GetFilteredCount()

	this.Data["filteredCount"] = filteredCount
	this.Data["illegalCount"] = allCount - filteredCount
}

// write json to client.
func (this *BaseController) writeJson(httpCode, dataCode int, data interface{}) {
	this.EnableRender = false

	this.Ctx.Output.Status = httpCode

	data = JSONData{
		Code:dataCode,
		Data:data,
	}

	this.Data["data"] = data
	this.ServeJson()
}