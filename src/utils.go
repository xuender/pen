package main
import (
  "path/filepath"
  "os"
  "crypto/md5"
  "encoding/hex"
  "bufio"
  "io"
  //"fmt"
  log "github.com/cihub/seelog"
)
func checkFiles() error{
  log.Info("开始创建文件摘要")
  t := [...]string{"css", ".js", "eot", "svg", "tff", "off", "eot", "tml"}
  fout,err := os.Create("base/checkData.go")
  defer fout.Close()
  if err != nil {
    return err
  }

  fout.WriteString("package base\n")
  fout.WriteString("func GetCheckData() map[string]string{\n")
  fout.WriteString("\tvData := make(map[string]string)\n")
  filepath.Walk(".",
  func(path string,fi os.FileInfo, err error) error {
    if (fi == nil) {
      return err
    }
    if fi.IsDir() {
      return nil
    }
    b := true
    for i := 0; i < len(t); i ++ {
      if(t[i] == path[len(path)-3:]){
        b=false
      }
    }
    if(b){
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
      fout.WriteString("\tvData[\""+path+"\"]=\""+md5v+"\"\n")
    }
    return nil
  })
  fout.WriteString("\treturn vData\n")
  fout.WriteString("}\n")
  log.Info("创建完毕")
  return nil
}
func main() {
  defer log.Flush()
  checkFiles()
}
