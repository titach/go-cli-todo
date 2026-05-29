package cmd

import (
	"fmt"
	"strconv"

	"github.com/titach/todo/db"
)

func RemoveTask(ids string) {
	id, notInt := strconv.Atoi(ids)
	db.CheckError(notInt)
	_, err := db.DB.Exec(
		"DELETE FROM todos WHERE id=$1;", id)
	db.CheckError(err)
	fmt.Println("Task clear.👋")
}