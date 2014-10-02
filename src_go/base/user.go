package base

import (
  "code.google.com/p/go.net/websocket"
  "errors"
  log "github.com/Sirupsen/logrus"
)

// 用户
type User struct {
  BaseObject
  // 性别
  Gender    int64
  // 昵称
  Nick      string
  Email     string
  Password  string  `json:"-"`
}

// 用户密码教研
func (u User) check(password string) bool {
  return u.Password == password
}

// 查找用户
func UserRead(nick string) (User, error) {
  var user User
  log.Debug(user.Id)
  db.Where("nick = ?", nick).First(&user)
  log.Debug(user.Id)
  if user.Id > 0 {
    return user, nil
  }
  return User{}, errors.New("用户没有找到")
}

// 获取全部用户
func userAllEvent(data *string, ws *websocket.Conn, session Session) {
  //TODO 权限认证
  var users []User
  db.Find(&users)
  send(ws, Code, 用户列表, users)
}

func init() {
  RegisterEvent(Code, 用户列表, userAllEvent)
  // 数据库初始化
  //db := GetDb()
  db.AutoMigrate(&User{})
  db.Model(&User{}).AddUniqueIndex("idx_user_nick", "nick")
  // 创建管理员
  var count int64
  db.Model(User{}).Count(&count)
  if count == 0 {
    e := User{
      Nick: "ender",
      Email: "xxx@xxx",
      Password: "d9b1d7db4cd6e70935368a1efb10e377",
    }
    log.WithFields(log.Fields{
      "id": e.Id,
    }).Debug("增加管理员")
    db.Create(&e)
    log.WithFields(log.Fields{
      "id": e.Id,
    }).Debug("增加管理员之后")
  }
}
