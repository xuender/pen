package utils
import (
  "testing"
)

func TestCommand(t *testing.T) {
  RegisterEvent("1", updateData)
  data := "old"
  Event("1", &data)
  if data != "new" {
    t.Errorf("字符串修改失败")
  }
}

func updateData(data *string){
  *data="new"
}
