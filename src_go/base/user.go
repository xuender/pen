package base

import (
	"../utils"
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"errors"
	log "github.com/Sirupsen/logrus"
	"time"
)

// 用户
type User struct {
	BaseObject
	// 性别
	Gender string `sql:type:char(1)`
	// 昵称
	Nick     string
	Email    string
	Password string `json:"-"`
}

// 用户密码教研
func (u *User) check(password string) bool {
	return u.Password == password
}

// 根据track计算token
func (u *User) getToken(track string) string {
	return utils.Md5(track + utils.Md5(time.Now().Format("2006-01-02")+u.Password))
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
	Send(ws, Code, 用户列表, users)
}

func updateUserEvent(data *string, ws *websocket.Conn, session Session) {
	//TODO 权限认证
	var u User
	json.Unmarshal([]byte(*data), &u)
	log.WithFields(log.Fields{
		"ID": u.Id,
	}).Debug("修改用户")
	if u.Id == 0 {
		e := db.Save(&u).Error
		if e == nil {
			//d.publish()
			Send(ws, Code, 修改用户, "ok")
		} else {
			log.Debug(e)
			Send(ws, Code, MSG, e)
		}
	} else {
		var o User
		db.First(&o, u.Id)
		o.Email = u.Email
		o.Gender = u.Gender
		e := db.Save(&o).Error
		if e == nil {
			//o.publish()
			Send(ws, Code, 修改用户, "ok")
		} else {
			log.Debug(e)
			Send(ws, Code, MSG, e)
		}
	}
}

// 初始化
func init() {
	RegisterEvent(Code, 用户列表, userAllEvent)
	RegisterEvent(Code, 修改用户, updateUserEvent)
	// 数据库初始化
	if db.CreateTable(&User{}).Error == nil {
		db.Model(&User{}).AddUniqueIndex("idx_user_nick", "nick")
		// 创建管理员
		e := User{
			Nick:     "ender",
			Gender:   "M",
			Email:    "xuender@gmail.com",
			Password: "40b0dada4577cd2a27d93ee392fa9a4f",
		}
		log.WithFields(log.Fields{
			"id": e.Id,
		}).Debug("增加管理员")
		db.Create(&e)
		db.Save(&Dict{Type: "type", Code: "gender", Title: "性别"})
		db.Save(&Dict{Type: "gender", Code: "M", Title: "男"})
		db.Save(&Dict{Type: "gender", Code: "F", Title: "女"})
	} else {
		db.AutoMigrate(&User{})
	}
}
