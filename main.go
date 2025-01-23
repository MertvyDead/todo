package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func help() {
	fmt.Println("Команды:\nadd \"Текст задачи\" - Добавить задачу,\nlist - Список задач,\ndone <ID> - отметить выполненную задачу,\nremove <ID> - удалить задачу,\nclear - удалить все задачи\nhelp - вывести это сообщение\nexit - выход из программы")
}

func main() {
	help()
	var (
		storage TaskManager = &TaskStorage{}
		scanner             = bufio.NewScanner(os.Stdin)
	)
	for {

		fmt.Println("\n>>> ")
		scanner.Scan()
		input := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Printf("Ошибка чтения: %v\n", err)
			break
		}
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}
		command := args[0]

		switch command {
		case "add":
			if len(args) < 2 {
				fmt.Println("Требуется title")
				continue
			}
			title := strings.Join(args[1:], " ")
			storage.AddTask(title)

		case "list":
			storage.ListTasks()

		case "done":
			if len(args) < 2 {
				fmt.Println("Требуется ID")
				continue
			}
			ID, err := strconv.Atoi(args[1])
			if err != nil || ID < 1 {

				fmt.Println("Неправильный ID")
				continue
			}
			err = storage.MarkDone(ID)
			if err != nil {

				fmt.Println(err)
				continue
			}
		case "remove":
			if len(args) < 2 {
				fmt.Println("Требуется ID")
				continue
			}
			ID, err := strconv.Atoi(args[1])

			if err != nil || ID < 1 {

				fmt.Println("Неправильный ID")
				continue
			}
			err = storage.DeleteTask(ID)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "clear":
			storage.Clear()
		case "help":
			help()
		case "exit":
			return
		default:
			fmt.Println("Неизвестная команда")
			help()
		}
	}
}
