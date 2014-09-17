package main

import (
  "github.com/eaigner/hood"
)

type Group struct {
  Id      hood.Id
  Name string
}
func (m *M) CreateGroup_1410942668_Up(hd *hood.Hood) {
  hd.CreateTable(&Group{})
}

func (m *M) CreateGroup_1410942668_Down(hd *hood.Hood) {
  hd.DropTable(&Group{})
}
