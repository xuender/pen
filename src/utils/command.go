package utils

type Command interface{}

var commands = make(map[string]Command)

func RegisterCommand(command string, Command){

}
