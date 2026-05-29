package cmd

import (
	"fmt"
	"time"

	"github.com/titach/todo/db"
)

func Createtodo(task string) {
	_, err := db.DB.Exec(
		"INSERT INTO todos(title, completed, created_at)VALUES ($1, $2, $3);", 
		task, false, time.Now().Format("2006-01-02"))

	db.CheckError(err)
	fmt.Println("✎ᝰ. Done!📝")
}