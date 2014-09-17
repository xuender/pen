package main

import (
	db "./db"
	log "github.com/cihub/seelog"
)

func main() {
	defer log.Flush()
	log.Info("启动")
	//pen.RunWeb()
	log.Info("关闭")
}
