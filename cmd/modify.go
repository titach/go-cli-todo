package cmd

import (
	"fmt"
	"strconv"

	"github.com/titach/todo/db"
)

func Updatetodo(ids string, newTask string) {
	id, notInt := strconv.Atoi(ids)
	db.CheckError(notInt)
	_, err := db.DB.Exec(
		"UPDATE todos SET title=$1 WHERE id=$2;",newTask,id)
	db.CheckError(err)
	fmt.Println("✎ᝰ. Done!📝")
}