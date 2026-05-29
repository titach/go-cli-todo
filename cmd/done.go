package cmd

import (
	"fmt"
	"strconv"

	"github.com/titach/todo/db"
)

func Donetask(ids string) {
	id, notInt := strconv.Atoi(ids)
	db.CheckError(notInt)
	_, err := db.DB.Exec(
		"UPDATE todos SET completed=$1 WHERE id=$2;", true, id)
	db.CheckError(err)
	fmt.Println("✎ᝰ. Done!📝")
}