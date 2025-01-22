package main

import (
	"fmt"
	"os"
	"strconv"
)

func help() {
	fmt.Println("Команды:\nadd \"Текст задачи\" - Добавить задачу,\nlist - Список задач,\ndone <ID> - отметить выполненную задачу,\nremove <ID> - удалить задачу,\nclear - удалить все задачи\nhelp - вывести это сообщение")
}

func main() {
	var storage TaskManager = &Storage{}

	if len(os.Args) < 2 {
		help()
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Требуется title")
			return
		}
		title := os.Args[2]
		storage.AddTask(title)
	case "list":
		storage.ListTasks()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Требуется ID")
			return
		}
		ID, err := strconv.Atoi(os.Args[2])
		if err != nil || ID < 1 {

			fmt.Println("Неправильный ID")
			return
		}
		err = storage.MarkDone(ID)
		if err != nil {

			fmt.Println(err)
			return
		}
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Требуется ID")
			return
		}
		ID, err := strconv.Atoi(os.Args[2])

		if err != nil || ID < 1 {

			fmt.Println("Неправильный ID")
			return
		}
		err = storage.DeleteTask(ID)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "clear":
		storage.Clear()
	case "help":
		help()
	default:
		fmt.Println("Неизвестная команда")
		help()
	}
}
