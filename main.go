package main

import (
	"fmt"
	"log"
	"os"

	"github.com/titach/todo/cmd"
	"github.com/titach/todo/db"
)

var arguments = os.Args
var argumentsLenght = len(arguments)

const (
	ADD = "add"
	LIST = "list"
	MODIFY = "modify"
	CHECK = "check"
	DONE = "done"
	REMOVE = "remove"
)

func handleCommand(num int) {
	if argumentsLenght != num {
		log.Fatal("[Invalid command arguments!]")
	}
}

func handleCommands() {
	if (argumentsLenght > 4 || argumentsLenght == 1) {
		log.Fatal("[Invalid arguments!]")
	}
	switch arguments[1] {
	case ADD:
		handleCommand(3)
	case LIST:
		handleCommand(2)
	case MODIFY:
		handleCommand(4)
	case DONE:
		handleCommand(3)
	case REMOVE:
		handleCommand(3)
	default:
		log.Fatal("[Invalid command!]")
	}
}

func main() {

	handleCommands()
	db.Connect_db()
	fmt.Println(cmd.CYAN+"═════════════")
	fmt.Println("CLI TODO APP")
	fmt.Println("═════════════"+cmd.RESET)
	do_command()
}

func do_command() {
	switch arguments[1] {
	case ADD:
		cmd.Createtodo(arguments[2])
	case LIST:
		cmd.GetAllTodos()
	case MODIFY:
		cmd.Updatetodo(arguments[2], arguments[3])
	case DONE:
		cmd.Donetask(arguments[2])
	case REMOVE:
		cmd.RemoveTask(arguments[2])
	}
}