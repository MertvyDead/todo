package internal

import (
	"fmt"
)

type TaskManager interface {
	AddTask(title string)
	MarkDone(ID int)
	ListTasks()
	DeleteTask(ID int)
	Clear()
}

type Storage struct {
	tasks []Task
}

func (s *Storage) AddTask(title string) {
	task := Task{ID: len(s.tasks) + 1, Title: title, IsDone: false}
	s.tasks = append(s.tasks, task)
	fmt.Printf("Задача \"%s\" добавлена", title)
}

func (s *Storage) ListTasks() error {
	if len(s.tasks) == 0 {

		return fmt.Errorf("Список задач пуст")
	}
	for _, task := range s.tasks {
		switch {
		case task.IsDone == false:
			fmt.Printf("%v. [ ] %s", task.ID, task.Title)
		case task.IsDone == true:
			fmt.Printf("%v. [X] %s", task.ID, task.Title)
		}

	}
	return nil
}

func (s *Storage) MarkDone(ID int) error {
	if ID <= 0 || ID > len(s.tasks)+1 {
		return fmt.Errorf("Неверный ID")
	}

	task := s.tasks[ID-1]
	if task.IsDone == true {
		fmt.Printf("Задача \"%s\" уже выполнена", task.Title)
	} else {
		task.done()
		fmt.Printf("Теперь задача \"%s\" выполнена", task.Title)
	}
	return nil
}
func (s *Storage) DeleteTask(ID int) error {
	if ID <= 0 || ID > len(s.tasks)+1 {
		return fmt.Errorf("Неверный ID")
	}

	return nil
}
func (s *Storage) Clear() {
	s.tasks = []Task{}
}
