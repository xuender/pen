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
func FindUser(name string) (User, error){
  if name == "ender" {
    return  User{Id:2, Nick:"xcy", Email:"xcy@gmail.com", Password:"123"}, nil
  }
  //for _, v := range users{
  //  if v.Nick == name{
  //    return v, nil
  //  }
  //}
  return User{}, errors.New("用户没有找到")
}
