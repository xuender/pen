package base

import (
	"../utils"
	"errors"
	log "github.com/Sirupsen/logrus"
	"os"
	"path/filepath"
)

func CheckFiles() error {
	// 文件检查
	ret := 0
	log.Info("文件完整性检查")
	vData := GetCheckData()
	filepath.Walk(".",
		func(path string, fi os.FileInfo, err error) error {
			if fi == nil {
				return err
			}
			if fi.IsDir() {
				return nil
			}
			v, ok := vData[path]
			if !ok {
				return nil
			}
			md5v, err := utils.Md5File(path)
			if err == nil {
				log.WithFields(log.Fields{
					"path": path,
				}).Info("检查")
				if v != md5v {
					log.WithFields(log.Fields{
						"path": path,
					}).Error("检查不匹配")
					ret++
				}
			}
			return nil
		})
	if ret > 0 {
		return errors.New("文件检查失败")
	}
	return nil
}
