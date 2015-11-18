package main

import (
	_ "github.com/gitwillsky/btsearcher.web/models"
	_ "github.com/gitwillsky/btsearcher.web/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
)

func init() {
	beego.AddFuncMap("size", Size)
}

// 转换字节数为对应大小格式
func Size(length uint64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	var mod float64
	mod = 1024.0
	size := float64(length)
	i := 0
	for size >= mod {
		size /= mod
		i++
	}

	return fmt.Sprintf("%.2f %s", size, units[i])
}

func main() {

	if beego.RunMode != "dev" {
		beego.BeeLogger.SetLevel(logs.LevelWarning)
		beego.BeeLogger.DelLogger("console")
		beego.SetLogger("file", `{"filename":"logs/app.log"}`)
	} else {
		beego.BeeLogger.SetLevel(logs.LevelDebug)
	}

	beego.Run()
}
