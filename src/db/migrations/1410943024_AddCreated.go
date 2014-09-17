package main

import (
	"github.com/eaigner/hood"
)

func (m *M) AddCreated_1410943024_Up(hd *hood.Hood) {
  hd.AddColumns("group", struct {
    Created hood.Created
    Updated hood.Updated
  }{})
}

func (m *M) AddCreated_1410943024_Down(hd *hood.Hood) {
  hd.RemoveColumns("group", struct {
    Created hood.Created
    Updated hood.Updated
  }{})
}
