package main

import "github.com/jinzhu/gorm"

type Task struct {
	Id          uint
	Name        string
	IsCompleted bool
	Tags        StringSlice
}

func GetTasksByTag(db *gorm.DB, tag string) ([]Task, error) {
	tasks := []Task{}
	result := db.Find(&tasks, "? = any(tags)", tag)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func CreateTask(db *gorm.DB, name string, tags []string) (uint, error) {
	newTask := &Task{Name: name, Tags: tags}
	result := db.Create(newTask)
	if result.Error != nil {
		return 0, result.Error
	}
	return newTask.Id, nil
}
