package base

import "errors"
// 用户
type User struct{
  Id        int64
  Nick      string
  Email     string
  Password  string
}
// 用户密码教研
func (u User) check(password string) bool{
  return u.Password == password
}

// TODO 使用数据库替换
// 查找用户
func FindUser(nick string) (User, error){
  if nick == "ender" {
    return  User{Id:2, Nick:"ender", Email:"xxx@xxx", Password:"d9b1d7db4cd6e70935368a1efb10e377"}, nil
  }
  return User{}, errors.New("用户没有找到")
}
