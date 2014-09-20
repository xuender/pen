package main

import (
	"github.com/eaigner/hood"
)
type User struct {
  Id        hood.Id
  Nick      string
  Email     string
  Mobile    string
  Password  string
  IsStaff   bool
  IsActive  bool
  Created   hood.Created
  Updated   hood.Updated
}
func (m *M) AddUser_1410943885_Up(hd *hood.Hood) {
  hd.CreateTable(&User{})
  hd.CreateIndex("user", "i_user_nick", true, "nick")
  hd.CreateIndex("user", "i_user_email", true, "email")
  hd.CreateIndex("user", "i_user_mobile", true, "mobile")
}

func (m *M) AddUser_1410943885_Down(hd *hood.Hood) {
  hd.DropTable(&User{})
}
