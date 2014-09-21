package base

import (
  "../utils"
  "code.google.com/p/go.net/websocket"
  log "github.com/cihub/seelog"
  "encoding/json"
)

// 消息内容
type WsMessage struct {
  Event   string
  Data    string
  Name    string
  Md5     string
}
// ws消息
type WsData struct {
  Message WsMessage
}
// 在线会话
type Session struct{
  IsLogin   bool
  Nick      string
  User      User
}

// 在线用户
var onlines = make(map[*websocket.Conn]Session)
// 接收ws消息
func WsHandler(ws *websocket.Conn) {
  var err error
  for {
    var reply WsData
    if err = websocket.JSON.Receive(ws, &reply); err != nil {
      delete(onlines, ws)
      log.Error("关闭连接, 在线用户:%d", len(onlines))
      // TODO 删除用户事件
      break
    }
    log.Debugf("接收信息: %s", reply)
    session, ok := onlines[ws]
    if !ok {
      user, err := FindUser(reply.Message.Name)
      session.IsLogin = false
      session.Nick = "来宾"
      if err == nil {
        session.IsLogin = true
        session.Nick = user.Nick
        session.User = user
      }
      // TODO 消息身份认证
      onlines[ws] = session
      // TODO 增加用户通知所有客户端，需要调用观察者
      if err = websocket.JSON.Send(ws, ok); err != nil {
        log.Error("不能发送消息到客户端")
        break
      }
    }
    // TODO 消息身份认证
    log.Debugf("收到消息,来自:%s", session.Nick)
    // 消息处理
    utils.Event(reply.Message.Event, &reply.Message.Data)
    //user, ok := onlines[ws]
    //if !ok {
    //  log.Debugf("初始化, 在线用户:%d", len(onlines))
    //  user = reply.Message.Source
    //  onlines[ws] = user
    //}
    //switch reply.Message.Command {
    //case "login":
    //  loginHandler(ws, user, reply.Message)
    //case "logout":
    //  logoutHandler(ws, user, reply.Message)
    //case "chat":
    //  chatHandler(ws, user, reply.Message)
    //case "init":
    //  initHandler(ws, user, reply.Message)
    //}
  }
}
// 登录信息
type LoginData struct {
  Nick    string  `json:"nick"`
  Token   string  `json:"token"`
}
// 登录事件
func loginEvent(data *string){
  log.Info(*data)
  var ld LoginData
  if err := json.Unmarshal([]byte(*data), &ld); err == nil {
    log.Debugf("[ %s ] 开始登录", ld.Nick)
  }
}
// 初始化
func init(){
  utils.RegisterEvent("base.login", loginEvent)
}
