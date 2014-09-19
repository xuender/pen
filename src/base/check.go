package base
import (
  log "github.com/cihub/seelog"
  "os"
  "path/filepath"
  "crypto/md5"
  "encoding/hex"
  "bufio"
  "io"
  "errors"
)

func CheckFiles() error{
  // 文件检查
  ret := 0
  log.Info("文件完整性检查")
  vData := GetCheckData()
  filepath.Walk(".",
  func(path string,fi os.FileInfo, err error) error {
    if (fi == nil) {
      return err
    }
    if fi.IsDir() {
      return nil
    }
    v, ok := vData[path]
    if(!ok){
      return nil
    }
    h := md5.New()
    f, err := os.Open(path)//打开文件
    defer f.Close() //打开文件出错处理
    if nil == err {
      buff := bufio.NewReader(f) //读入缓存
      for {
        line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
        if err != nil || io.EOF == err {
          break
        }
        h.Write([]byte(line))
      }
      md5v := hex.EncodeToString(h.Sum(nil))
      log.Debugf("check:%s",path)
      if(v!=md5v){
        log.Errorf("文件 %s 不匹配", path)
        ret ++
      }
    }
    return nil
  })
  if(ret > 0){
    return errors.New("文件检查失败")
  }
  return nil
}
