package base

import (
  "../utils"
  "code.google.com/p/go.net/websocket"
  log "github.com/cihub/seelog"
)

type WsMessage struct {
  Event   string
  Data    string
}
// ws消息
type WsData struct {
  Message WsMessage
}

// 接收ws请求
func WsHandler(ws *websocket.Conn) {
  var err error
  for {
    var reply WsData
    if err = websocket.JSON.Receive(ws, &reply); err != nil {
      //delete(onlines, ws)
      log.Error("关闭连接, 在线用户:%d", 0) // len(onlines))
      //usersHandler()
      break
    }
    log.Debugf("接收信息: %s", reply)
    // TODO 消息身份认证
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
