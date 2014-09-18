package main

import (
  //db "./db"
  log "github.com/cihub/seelog"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/martini-contrib/gzip"
)

func main() {
  defer log.Flush()
  log.Info("启动")
  m := martini.Classic()
  // gzip支持
  m.Use(gzip.All())
  // 模板支持
  m.Use(render.Renderer())
  //m.Get("/", func() string {
  //  return "wwww"
  //})
  m.NotFound(func(r render.Render){
    r.HTML(404, "404", nil)
  })
  m.Run()
  log.Info("关闭")
}
