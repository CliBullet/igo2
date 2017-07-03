package main

/* This sample assumes that you have a local postgres database with the name `mytasks`
and a table, 'tasks', with the following schema

	CREATE TABLE tasks (
		id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    is_completed BOOL NOT NULL,
    tags VARCHAR(10)[]
	)

	INSERT INTO tasks(name,is_completed,tags) VALUES('buy milk',false,'{home,delegate}');

*/

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := gorm.Open("postgres",
		`host=localhost 
			user=postgres password=test
			dbname=mytasks 
			sslmode=disable`)
	panicOnError(err)
	defer db.Close()

	id, err := CreateTask(db, "test 123", []string{"personal", "test"})
	panicOnError(err)
	fmt.Printf("Task %d has been created\n", id)

	tasks, err := GetAllTasks(db)
	panicOnError(err)
	fmt.Println(tasks)
}
