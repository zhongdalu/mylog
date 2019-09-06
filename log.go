//@Auth:zdl
package mylog

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/util/gconv"
)

func init() {
	var debug = gconv.Bool(g.Config().Get("log.isDebug"))
	glog.SetDebug(debug)
	glog.SetAsync(true)
	err := glog.SetPath(g.Config().GetString("log.path"))
	if err != nil {
		panic(err)
	}
	style := g.Config().Get("log.style")
	glog.SetFile(gconv.String(style))
}

func getLogger() *glog.Logger {
	return glog.Cat(gtime.Date())
}

func Debug(v ...interface{}) {
	getLogger().Async().Debug(v...)
}

func Println(v ...interface{}) {
	getLogger().Async().Println(v...)
}

func Fatal(v ...interface{}) {
	getLogger().Fatal(v...)
}

func Error(v ...interface{}) {
	getLogger().Async().Error(v...)
}
