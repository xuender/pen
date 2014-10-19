package main

import (
	//db "./db"
	"./base"
	_ "./class"
	"code.google.com/p/go.net/websocket"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"net/http"
	//"./utils"
)

func main() {
	log.Info("启动")
	//读取配置文件
	base.BaseConfig.Read("config.json")
	// 初始化数据库
	base.InitDb()
	// 检查文件完整性
	//err := base.CheckFiles()
	//if(err!=nil){
	//  log.Error(err)
	//  return
	//}
	m := martini.Classic()
	// gzip支持
	m.Use(gzip.All())
	// 模板支持
	m.Use(render.Renderer())
	//
	m.Get("/meta.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write([]byte(base.GetMetaJs()))
	})
	// websocket 支持
	m.Get("/ws", websocket.Handler(base.WsHandler).ServeHTTP)
	m.NotFound(func(r render.Render) {
		r.HTML(404, "404", nil)
	})
	log.Info(fmt.Sprintf("访问地址 http://localhost:%d", base.BaseConfig.Web.Port))
	// 端口号
	http.ListenAndServe(fmt.Sprintf(":%d", base.BaseConfig.Web.Port), m)
	// m.Run()
	log.Info("退出")
}
