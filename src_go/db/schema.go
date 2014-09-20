package db

import (
	"github.com/eaigner/hood"
)

type Group struct {
	Id      hood.Id
	Name    string
	Created hood.Created
	Updated hood.Updated
}

type User struct {
	Id       hood.Id
	Nick     string
	Email    string
	Mobile   string
	Password string
	IsStaff  bool
	IsActive bool
	Created  hood.Created
	Updated  hood.Updated
}

func (table *User) Indexes(indexes *hood.Indexes) {
	indexes.AddUnique("i_nick", "nick")
	indexes.AddUnique("i_email", "email")
	indexes.AddUnique("i_mobile", "mobile")
}
