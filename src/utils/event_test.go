package utils
import (
  "testing"
)

func TestCommand(t *testing.T) {
  data := "old"
  Event("1", &data)
  if data != "old" {
    t.Errorf("字符串不应被修改")
  }
  RegisterEvent("1", updateData)
  data = "old"
  Event("1", &data)
  if data != "new" {
    t.Errorf("字符串没有被修改")
  }
}

func updateData(data *string){
  *data="new"
}
