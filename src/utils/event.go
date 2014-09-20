package utils
import (
  "reflect"
)

var commands = make(map[string]func(*string))

// 注册事件
func RegisterEvent(event string,command func(*string)){
  if reflect.TypeOf(command).Kind() != reflect.Func {
    panic("command must be a callable func")
    return
  }
  commands[event] = command
}

// 事件触发
func Event(event string,data *string){
  command, ok := commands[event]
  if ok{
    command(data)
  }
}
