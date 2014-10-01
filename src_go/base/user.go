package base

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"errors"
	"fmt"
  "time"
	log "github.com/Sirupsen/logrus"
)

// 用户
type User struct {
  Id        int64
  // 性别
  Gender    int64
  // 昵称
	Nick      string
	Email     string
	Password  string
  Created   time.Time
  Updated   time.Time
}

// 用户密码教研
func (u User) check(password string) bool {
	return u.Password == password
}

var ender = User{Id: 2, Nick: "ender", Email: "xxx@xxx", Password: "d9b1d7db4cd6e70935368a1efb10e377"}

// 查找用户 TODO 使用数据库替换
func UserRead(nick string) (User, error) {
	if nick == "ender" {
		return ender, nil
	}
	return User{}, errors.New("用户没有找到")
}

// 查找所有用户
func UserAll() []User {
	var users []User
	//ender.Password = ""
	users = append(users, ender)
	for i := 0; i < 200; i++ {
		users = append(users, User{Id: int64(3 + i), Nick: fmt.Sprintf("user:%d", i), Email: "xxx@xxx", Password: ""})
	}
	return users
}

// 获取全部用户
func userAllEvent(data *string, ws *websocket.Conn, session Session) {
	//TODO 权限认证
	users := UserAll()
	s, err := json.Marshal(users)
	if err != nil {
		log.WithFields(log.Fields{
			"users": users,
		}).Error("JSON编码错误")
	} else {
		send(ws, Code, 用户列表, string(s))
	}
}

func init() {
	RegisterEvent(Code, 用户列表, userAllEvent)
  db := GetDb()
  db.AutoMigrate(&User{})
}
