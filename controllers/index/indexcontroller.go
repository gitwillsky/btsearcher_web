package index
import (
	"github.com/astaxie/beego/utils/pagination"
	"github.com/gitwillsky/btsearcher.web/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type IndexController struct {
	BaseController
}

// index page
// @router / [get]
func (this *IndexController) Index() {
	this.TplNames = "front/index.html"

	this.Data["hots"], _ = models.Search("", false, 4, 0, "-hot")

	this.Render()
}

// search list
// @router /list.html [get]
func (this *IndexController) List() {
	var err error

	this.TplNames = "front/list.html"

	numPerPage := 10
	keywords := this.GetString("keywords")
	all := strings.HasPrefix(keywords, "all:")
	if all {
		keywords = strings.Split(keywords, "all:")[1]
	}

	maxCount, _ := models.SearchCount(keywords, all)
	this.Data["count"] = maxCount

	paginator := pagination.SetPaginator(this.Ctx, numPerPage, maxCount)
	data, err := models.Search(keywords, all, numPerPage, paginator.Offset(), "-create_at")
	if err != nil {
		beego.Error(err)
	}

	for _, torrent := range data {
		if len(torrent.FileList) > 5 {
			torrent.FileList = torrent.FileList[:5]
		}
	}

	this.Data["torrents"] = data
	this.Render()
}

// view
// @router /view/:id:int.html [get]
func (this *IndexController) View() {
	this.TplNames = "front/view.html"

	id, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 0)

	torrent, err := models.View(id)
	if err != nil {
		beego.Error(err)
	}

	this.Data["torrent"] = torrent

	this.Render()
}