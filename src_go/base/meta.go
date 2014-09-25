package base


// 模块元信息
type Meta struct {
  Name          string
  Code          string
  Description   string
  Action        []int
}

var metas []Meta

// 注册元信息
func RegisterMeta(meta Meta){
  metas = append(metas, meta)
}
