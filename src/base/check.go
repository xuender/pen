package base
import (
  log "github.com/cihub/seelog"
  "os"
  "path/filepath"
  "crypto/md5"
  "encoding/hex"
  "bufio"
  "io"
  "fmt"
  "errors"
)

func CheckFiles(check bool) error{
  // 文件检查
  ret := 0
  vData := make(map[string]string)
  if(check){
    log.Info("文件完整性检查")
    vData["index.html"]="d3c911705b6cf8a0872bd69f30332de0"
    vData["public/css/bootstrap-theme.css.map"]="d41d8cd98f00b204e9800998ecf8427e"
    vData["public/css/bootstrap-theme.min.css"]="1fa9f46684e88960d291916cc9d4119c"
    vData["public/css/bootstrap.css.map"]="d41d8cd98f00b204e9800998ecf8427e"
    vData["public/css/bootstrap.min.css"]="1fa9f46684e88960d291916cc9d4119c"
    vData["public/css/font-awesome.min.css"]="354b1b9dd22d3bbd86f0d67255f45068"
    vData["public/fonts/FontAwesome.otf"]="bff5682c9f7c82e93405136aed960082"
    vData["public/fonts/fontawesome-webfont.eot"]="e84b691c1f7be4eb0d11fd979fa1a76a"
    vData["public/fonts/fontawesome-webfont.svg"]="9693cf85a804fd390e059a334bf5bf90"
    vData["public/fonts/fontawesome-webfont.ttf"]="5c42afb75b45144c023a82d8d50c6fd0"
    vData["public/fonts/fontawesome-webfont.woff"]="b1da2c3c3001b34dbcbac4209732a71b"
    vData["public/fonts/glyphicons-halflings-regular.eot"]="12e601026da631ce23f773bba9f1cc79"
    vData["public/fonts/glyphicons-halflings-regular.svg"]="b1224c183966afebeddf7079a5f42732"
    vData["public/fonts/glyphicons-halflings-regular.ttf"]="9c7ca70a72d8e030851658dd7e13d60f"
    vData["public/fonts/glyphicons-halflings-regular.woff"]="b9bb2b8ac3cbd0f14214d8c511c05d46"
    vData["public/index.html"]="d3c911705b6cf8a0872bd69f30332de0"
    vData["public/js/angular-file-upload-html5-shim.min.js"]="cd5a9beae042d50cbcbeaf4b96432e52"
    vData["public/js/angular-file-upload.min.js"]="cd5a9beae042d50cbcbeaf4b96432e52"
    vData["public/js/angular-local-storage.min.js"]="d41d8cd98f00b204e9800998ecf8427e"
    vData["public/js/angular.js"]="9cbb23e932f43c79ea66eb97291420ac"
    vData["public/js/angular.min.js"]="48eeee3bd2b189eb98b3d33e561ad614"
    vData["public/js/angular.min.js.map"]="e227e8367d48bf14aa38e5188d5f1494"
    vData["public/js/bootstrap.min.js"]="a6039569de87df380062f07816250e2e"
    vData["public/js/hotkey.min.js"]="d41d8cd98f00b204e9800998ecf8427e"
    vData["public/js/jquery.ba-resize.min.js"]="735da2a7c5ed19a59f7eb9b72c1ac197"
    vData["public/js/jquery.min.js"]="e40ec2161fe7993196f23c8a07346306"
    vData["public/js/jquery.min.map"]="d41d8cd98f00b204e9800998ecf8427e"
    vData["public/js/ngSocket.js"]="6febaba7868c5ea179c7ba4e00be8e17"
    vData["public/js/textAngular-sanitize.min.js"]="d41d8cd98f00b204e9800998ecf8427e"
    vData["public/js/textAngular.min.js"]="a074634b7aba905aa60e198f4970e45f"
    vData["public/js/ui-bootstrap-tpls.min.js"]="d49393c58ff7d4fb9b50f013ed3a6eb4"
    vData["templates/404.tmpl"]="b5cff082046093d7017a5dab6271a63b"
  }else{
    fmt.Println("vData := make(map[string]string)")
  }
  filepath.Walk(".",
  func(path string,fi os.FileInfo, err error) error {
    if (fi == nil) {
      return err
    }
    if fi.IsDir() {
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
      if(check){
        log.Infof("check:%s",path)
        v, ok := vData[path]
        if(ok && v!=md5v){
          log.Errorf("文件 %s 不匹配", path)
          ret ++
        }
      }else{
        fmt.Printf("vData[\"%s\"]=\"%s\"\n", path,md5v)
      }
    }
    return nil
  })
  if(ret > 0){
    return errors.New("文件检查失败")
  }
  return nil
}
