package main

import (
	//db "./db"
	"./base"
	"code.google.com/p/go.net/websocket"
	log "github.com/cihub/seelog"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
  "./utils"
)

func main() {
	defer log.Flush()
	log.Info("启动")
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
	// websocket 支持
	m.Get("/ws", websocket.Handler(base.WsHandler).ServeHTTP)
	m.NotFound(func(r render.Render) {
		r.HTML(404, "404", nil)
	})
  // 注册接收测试消息
  utils.RegisterEvent("test", testEvent)
	m.Run()
	log.Info("退出")
}
func testEvent(data *string){
  log.Info(*data)
}