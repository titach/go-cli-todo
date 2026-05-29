package cmd

import (
	"fmt"

	"github.com/titach/todo/db"
)

func GetAllTodos() {
	rows, err := db.DB.Query("SELECT id, title, completed, created_at FROM todos;")
	db.CheckError(err)
	defer rows.Close()

	var items []Todos
	for rows.Next() {
		var item Todos
		err = rows.Scan(&item.ID, &item.Title, &item.Check, &item.Date)
		db.CheckError(err)
		items = append(items, item)
	}
	err = rows.Err()
	db.CheckError(err)
	if len(items) == 0 {
		fmt.Println("no data📂")
	} else {
		tablelist(items)
	}
}

func tablelist(list []Todos) {
	fmt.Print(BLUE+"ID ")
	fmt.Print(YELLOW+"title ")
	fmt.Print(GREEN+"check ")
	fmt.Println(MAGENTA+"date"+RESET)
	fmt.Println("───────────────────")

	for _, row := range list {
		fmt.Print(BLUE,row.ID," ")
		fmt.Print(YELLOW,row.Title," ",GREEN)
		if row.Check == false {
			fmt.Print("⛶ ")
		} else {
			fmt.Print("✓ ")
		}
		fmt.Println(MAGENTA,row.Date.Format("2006-01-02"),RESET)
	}
	fmt.Println("───────────────────")
}