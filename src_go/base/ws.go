package base

import (
  "reflect"
  "../utils"
  "code.google.com/p/go.net/websocket"
  "code.google.com/p/go-uuid/uuid"
  log "github.com/cihub/seelog"
  "encoding/json"
  "unsafe"
  "fmt"
  "time"
  "gopkg.in/fatih/set.v0"
)

// 消息内容
type WsMessage struct {
  // 消息事件
  Event   string    `json:"event"`
  // 消息体
  Data    string    `json:"data"`
  // 身份认证令牌
  Token   string    `json:"token"`
  // 管道
  Tract   string    `json:"tract"`
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
  // 管道
  Tract   string
  // 身份认证令牌
  Token   string
}

// 在线用户
var onlines = make(map[*websocket.Conn]Session)
// 接收ws消息
func WsHandler(ws *websocket.Conn) {
  var err error
  for {
    var reply WsData
    if err = websocket.JSON.Receive(ws, &reply); err != nil {
      // 取消一切观察
      RemoveOb(ws)
      delete(onlines, ws)
      updateCount()
      log.Error("关闭连接, 在线用户:%d", len(onlines))
      // TODO 删除用户事件
      break
    }
    log.Debugf("接收信息: %s", reply)
    session, ok := onlines[ws]
    if !ok {
      //user, err := FindUser(reply.Message.Name)
      session.IsLogin = false
      session.Nick = "来宾"
      session.Tract = fmt.Sprintf("%d", unsafe.Pointer(ws))
      session.Token = uuid.NewUUID().String()
      onlines[ws] = session
      // 服务器要求客户端发送登陆信息
      send(ws, "base.login", "login")
      // 开始关注在线人数
      RegisterOb("base.count", ws)
    }else{
      log.Debugf("收到消息,来自:%s Token:%s message.Token:%s", session.Nick, session.Token, reply.Message.Token)
      if session.Token == reply.Message.Token {
        // 消息处理
        Event(reply.Message.Event, &reply.Message.Data, ws)
      }else{
        Event("base.login", &reply.Message.Data, ws)
      }
    }
    //user, ok := onlines[ws]
    //if !ok {
    //  log.Debugf("初始化, 在线用户:%d", len(onlines))
    //  user = reply.Message.Source
    //  onlines[ws] = user
    //}
  }
}
// 发送消息
func send(ws *websocket.Conn,event string, data string){
  var m WsMessage
  m.Event = event
  m.Data = data
  session := onlines[ws]
  m.Tract = session.Tract
  s, err := json.Marshal(m);
  if err != nil {
    log.Errorf("JSON编码错误: %s", data)
  }
  log.Debugf("send: 发送的数据: %s",string(s))
  if err = websocket.JSON.Send(ws, string(s)); err != nil {
    log.Error("不能发送消息到客户端")
  }
}
// 登录信息
type LoginData struct {
  Nick    string  `json:"nick"`
  Token   string  `json:"token"`
}
// 登录事件
// 服务器保存[加密密码]=md5(md5(密码))，加密密码保存在数据库
// 用户本地存储保存[当天加密密码]=md5(日期+加密密码)
// 客户端访问服务器获取管道(Tract),生成[令牌(Token)]=md5(管道track+当天加密密码)
// 客户端将令牌Token保存在本地存储，每次访问提交令牌，直到令牌过期
func loginEvent(data *string, ws *websocket.Conn, session Session){
  log.Info(*data)
  var ld LoginData
  if err := json.Unmarshal([]byte(*data), &ld); err == nil {
    log.Debugf("[ %s ] 开始登录", ld.Nick)
    if user, e := FindUser(ld.Nick); e == nil{
      token := utils.Md5(session.Tract + utils.Md5(time.Now().Format("2006-01-02") + user.Password))
      log.Debugf("服务器token: %s, 客户端token: %s", token, ld.Token)
      if token == ld.Token {
        session.Nick = user.Nick
        session.Token = token
        session.User = user
        session.IsLogin = true
        onlines[ws] = session
        log.Debugf("ws.Token:%s, session.Token:%s", onlines[ws].Token, session.Token)
        updateCount()
      }else{
        send(ws, "base.login", "ERROR_PASSWORD")
      }
    }else{
      send(ws, "base.login", "ERROR_NICK")
    }
  }else{
    send(ws, "base.login", "ERROR_DATA")
  }
}
// 统计在线人数
func updateCount(){
  count := 0
  for _, session:= range onlines {
    if session.IsLogin {
      count++
    }
  }
  ObUpdate("base.count", "base.count", fmt.Sprintf("%d", count))
}
// 登出
func logoutEvent(data *string, ws *websocket.Conn, session Session){
  log.Debugf("用户登出: %s", session.User.Nick)
  session.IsLogin = false
  session.Token = uuid.NewUUID().String()
  session.Nick = "来宾"
  updateCount()
  //session.User = nil
  //TODO 人数统计修改
}
// 初始化
func init(){
  RegisterEvent("base.login", loginEvent)
  RegisterEvent("base.logout", logoutEvent)
}
// 交互观察者
var obmap = make(map[string]*set.Set)
// 取消观察者
func RemoveNameOb(name string, ws *websocket.Conn){
  s, ok := obmap[name]
  if ok{
    s.Remove(ws)
  }
}
// 取消观察者
func RemoveOb(ws *websocket.Conn){
  for _, s := range obmap{
    s.Remove(ws)
  }
}
// 注册观察者
func RegisterOb(name string, ws *websocket.Conn){
  s, ok := obmap[name]
  if ok{
    s.Add(ws)
  }else{
    s = set.New()
    s.Add(ws)
    obmap[name] = s
  }
}
// 消息分发
func ObUpdate(name string, event string, data string){
  s, ok := obmap[name]
  if ok {
    l := s.Size()
    items := s.List()
    for i:=0;i<l;i++{
      ws, ok := items[i].(*websocket.Conn)
      if ok{
        send(ws, event, data)
      }
    }
  }
}
// ws事件
var commands = make(map[string]func(*string, *websocket.Conn, Session))

// 注册事件
func RegisterEvent(event string,command func(*string, *websocket.Conn, Session)){
  if reflect.TypeOf(command).Kind() != reflect.Func {
    panic("command must be a callable func")
    return
  }
  commands[event] = command
}

// 事件触发
func Event(event string, data *string, ws *websocket.Conn){
  command, ok := commands[event]
  if ok{
    session, ok := onlines[ws]
    if ok{
      command(data, ws, session)
    }
  }
}
