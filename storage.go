package main

import (
	"fmt"
)

type TaskManager interface {
	AddTask(title string)
	MarkDone(ID int) error
	ListTasks() error
	DeleteTask(ID int) error
	Clear()
}

type TaskStorage struct {
	tasks []*Task
}

func (s *TaskStorage) AddTask(title string) {
	task := &Task{ID: len(s.tasks) + 1, Title: title, IsDone: false}
	s.tasks = append(s.tasks, task)
	fmt.Printf("Задача \"%s\" добавлена", title)
}

func (s *TaskStorage) ListTasks() error {
	if len(s.tasks) == 0 {
		return fmt.Errorf("cписок задач пуст")
	}
	for _, task := range s.tasks {
		status := ""
		switch task.IsDone {
		case false:
			status = "[ ]"
			fmt.Printf("%v. %s %s\n", task.ID, status, task.Title)
		case true:
			status = "[X]"
			fmt.Printf("%v. %s %s\n", task.ID, status, task.Title)
		}

	}
	return nil
}

func (s *TaskStorage) MarkDone(ID int) error {
	if ID <= 0 || ID > len(s.tasks)+1 {
		return fmt.Errorf("неверный ID")
	}
	task := s.tasks[ID-1]
	task.done()
	fmt.Printf("Теперь задача \"%s\" выполнена", task.Title)
	return nil
}

func (s *TaskStorage) DeleteTask(ID int) error {
	if ID <= 0 || ID > len(s.tasks)+1 {
		return fmt.Errorf("неверный ID")
	}
	s.tasks = append(s.tasks[:ID-1], s.tasks[ID:]...)
	for i, v := range s.tasks {
		v.newID(i + 1)

	}
	return nil
}
func (s *TaskStorage) Clear() {
	s.tasks = []*Task{}
}
