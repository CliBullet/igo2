package main

/* This sample assumes that you have a local postgres database with the name `mytasks`
and a table, 'tasks', with the following schema

	CREATE TABLE tasks (
		id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    is_completed BOOL NOT NULL,
    tags VARCHAR(10)[]
	)

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

	data := []struct {
		task string
		tags []string
	}{
		{"call mom", []string{"family"}},
		{"schedule meeting with the team", []string{"project-x"}},
		{"prepare for client demo", []string{"slides", "project-x"}},
		{"book ticket", []string{"travel", "delegate"}},
	}

	for _, d := range data {
		id, err := CreateTask(db, d.task, d.tags)
		panicOnError(err)
		fmt.Printf("(%d) - Task %s has been created\n", id, d.task)
	}

	tasks, err := GetTasksByTag(db, "project-x")
	panicOnError(err)
	fmt.Println(tasks)
}
