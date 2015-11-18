package models
import (
	"github.com/gitwillsky/btsearcher/models"
	"github.com/astaxie/beego/orm"
)


func SearchCount(name string, all bool) (int64, error) {
	o := orm.NewOrm()

	if all {
		return o.QueryTable("torrent").
		Filter("name__icontains", name).
		Count()
	}

	return o.QueryTable("torrent").
	Filter("name__icontains", name).
	Filter("enable", true).
	Count()
}

// search by limit offset orderby
func Search(name string, all bool, limit, offset int, order ...string) ([]*models.Torrent, error) {
	o := orm.NewOrm()

	var torrents []*models.Torrent

	if all {
		if _, err := o.QueryTable("torrent").
		Filter("name__icontains", name).
		Limit(limit, offset).
		OrderBy(order...).
		All(&torrents); err != nil {
			return nil, err
		}
	} else {
		if _, err := o.QueryTable("torrent").
		Filter("name__icontains", name).
		Filter("enable", true).
		Limit(limit, offset).
		OrderBy(order...).
		All(&torrents); err != nil {
			return nil, err
		}
	}

	// load relate
	for _, torrent := range torrents {
		o.LoadRelated(torrent, "FileList")
	}

	return torrents, nil
}

// view a single torrent
func View(id int64) (*models.Torrent, error) {
	o := orm.NewOrm()

	result := &models.Torrent{}

	err := o.QueryTable("torrent").Filter("id", id).One(result)
	if err != nil {
		return nil, err
	}

	// load relate
	o.LoadRelated(result, "FileList")

	return result, nil
}

// torrents count
func Count() (int64, error) {
	o := orm.NewOrm()

	return o.QueryTable("torrent").Count()
}

// 获取已经过滤的资源数目
func GetFilteredCount() (int64, error) {
	o := orm.NewOrm()

	return o.QueryTable("torrent").Filter("enable", true).Count()
}