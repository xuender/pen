package base

import (
	"code.google.com/p/go-uuid/uuid"
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/fatih/set.v0"
	"reflect"
	"unsafe"
)

// 类型消息
type TypeMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// 消息内容
type WsMessage struct {
	// 功能
	Code string `json:"code"`
	// 消息
	Event int `json:"event"`
	// 消息体
	Data string `json:"data"`
	// 身份认证令牌
	Token string `json:"token"`
	// 管道
	Tract string `json:"tract"`
}

// ws消息
type WsData struct {
	Message WsMessage
}

// 在线会话
type Session struct {
	IsLogin bool
	Nick    string
	User    User
	// 管道
	Tract string
	// 身份认证令牌
	Token string
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
			log.WithFields(log.Fields{
				"在线用户数": len(onlines),
			}).Info("管理连接")
			break
		}
		log.WithFields(log.Fields{
			"消息": reply,
		}).Debug("接收信息")
		session, ok := onlines[ws]
		if !ok {
			//user, err := UserRead(reply.Message.Name)
			session.IsLogin = false
			session.Nick = "来宾"
			session.Tract = fmt.Sprintf("%d", unsafe.Pointer(ws))
			session.Token = uuid.NewUUID().String()
			onlines[ws] = session
			// 服务器要求客户端发送登陆信息
			send(ws, Code, 登录, "login")
			// 开始关注在线人数
			RegisterOb(人数, ws)
		} else {
			log.WithFields(log.Fields{
				"用户":           session.Nick,
				"Token":        session.Token,
				"messageToken": reply.Message.Token,
			}).Debug("接收到信息")
			if session.Token == reply.Message.Token {
				// 消息处理
				Event(reply.Message.Code, reply.Message.Event, &reply.Message.Data, ws)
			} else {
				Event(Code, 登录, &reply.Message.Data, ws)
			}
		}
	}
}

// 发送消息
func send(ws *websocket.Conn, code string, event int, data interface{}) {
	var str string
	switch data.(type) {
	case string:
		str = data.(string)
	default:
		d, err := json.Marshal(data)
		if err != nil {
			log.WithFields(log.Fields{
				"JSON": data,
			}).Error("JSON编码错误")
			return
		}
		str = string(d)
	}
	var m WsMessage
	m.Code = code
	m.Event = event
	m.Data = str
	session := onlines[ws]
	m.Tract = session.Tract
	//s, err := json.Marshal(m)
	//if err != nil {
	//  log.WithFields(log.Fields{
	//    "JSON": data,
	//  }).Error("JSON编码错误")
	//}
	log.WithFields(log.Fields{
		"code":  code,
		"event": event,
		"数据":    m,
	}).Debug("send: 发送数据")
	if err := websocket.JSON.Send(ws, m); err != nil {
		log.Error("不能发送消息到客户端")
	}
}

// 登录信息
type LoginData struct {
	Nick  string `json:"nick"`
	Token string `json:"token"`
}

// 登录事件
// 服务器保存[加密密码]=md5(md5(用户名+密码))，加密密码保存在数据库
// 用户本地存储保存[当天加密密码]=md5(日期+加密密码)
// 客户端访问服务器获取管道(Tract),生成[令牌(Token)]=md5(管道track+当天加密密码)
// 客户端将令牌Token保存在JS变量，每次访问提交令牌，直到关闭连接
// 再次访问则使用[当天加密密码]重新生成本次连接的令牌
// 跨日则重新输入密码
func loginEvent(data *string, ws *websocket.Conn, session Session) {
	log.WithFields(log.Fields{
		"data": *data,
	}).Debug("loginEvent")
	var ld LoginData
	if err := json.Unmarshal([]byte(*data), &ld); err == nil {
		log.WithFields(log.Fields{
			"nick": ld.Nick,
		}).Debug("开始登录")
		if user, e := UserRead(ld.Nick); e == nil {
			token := user.getToken(session.Tract)
			log.WithFields(log.Fields{
				"服务器token": token,
				"客户端token": ld.Token,
			}).Debug("令牌验证")
			if token == ld.Token {
				session.Nick = user.Nick
				session.Token = token
				session.User = user
				session.IsLogin = true
				onlines[ws] = session
				log.WithFields(log.Fields{
					"ws.Token":      onlines[ws].Token,
					"session.Token": session.Token,
				}).Debug("令牌验证成功")
				send(ws, Code, 登录, "OK")
				updateCount()
			} else {
				send(ws, Code, 登录, "ERROR_PASSWORD")
			}
		} else {
			send(ws, Code, 登录, "ERROR_NICK")
		}
	} else {
		send(ws, Code, 登录, "ERROR_DATA")
	}
}

// 统计在线人数
func updateCount() {
	count := 0
	for _, session := range onlines {
		if session.IsLogin {
			count++
		}
	}
	ObUpdate(人数, Code, 人数, fmt.Sprintf("%d", count))
}

// 登出
func logoutEvent(data *string, ws *websocket.Conn, session Session) {
	log.WithFields(
		log.Fields{
			"用户": session.User.Nick,
		},
	).Debug("用户登出")
	session.IsLogin = false
	session.Token = uuid.NewUUID().String()
	session.Nick = "来宾"
	updateCount()
	//session.User = nil
	//TODO 人数统计修改
}

// 初始化
func init() {
	RegisterEvent(Code, 登录, loginEvent)
	RegisterEvent(Code, 登出, logoutEvent)
}

// 交互观察者
var obmap = make(map[int]*set.Set)

// 取消观察者
func RemoveNameOb(name int, ws *websocket.Conn) {
	s, ok := obmap[name]
	if ok {
		s.Remove(ws)
	}
}

// 取消观察者
func RemoveOb(ws *websocket.Conn) {
	for _, s := range obmap {
		s.Remove(ws)
	}
}

// 注册观察者
func RegisterOb(name int, ws *websocket.Conn) {
	s, ok := obmap[name]
	if ok {
		s.Add(ws)
	} else {
		s = set.New()
		s.Add(ws)
		obmap[name] = s
	}
}

// 消息分发
func ObUpdate(name int, code string, event int, data interface{}) {
	s, ok := obmap[name]
	if ok {
		l := s.Size()
		items := s.List()
		for i := 0; i < l; i++ {
			ws, ok := items[i].(*websocket.Conn)
			if ok {
				go send(ws, code, event, data)
			}
		}
	}
}

// ws事件
var commands = make(map[string]map[int]func(*string, *websocket.Conn, Session))

// 注册事件
func RegisterEvent(code string, event int, command func(*string, *websocket.Conn, Session)) {
	if reflect.TypeOf(command).Kind() != reflect.Func {
		panic("command must be a callable func")
		return
	}
	events, ok := commands[code]
	if !ok {
		events = make(map[int]func(*string, *websocket.Conn, Session))
		events[event] = command
		commands[code] = events
	} else {
		events[event] = command
	}
}

// 事件触发
func Event(code string, event int, data *string, ws *websocket.Conn) {
	log.WithFields(log.Fields{
		"code":  code,
		"event": event,
		"data":  *data,
	}).Debug("Event")
	events, ok := commands[code]
	log.WithFields(log.Fields{
		"ok": ok,
	}).Debug("命令查找")
	if ok {
		command, ok := events[event]
		if ok {
			session, ok := onlines[ws]
			if ok {
				command(data, ws, session)
			}
		}
	}
}
