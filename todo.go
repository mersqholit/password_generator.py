package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tasks []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTo-Do List:")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Remove Task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task: ")
			scanner.Scan()
			task := scanner.Text()
			tasks = append(tasks, task)
			fmt.Println("Task added!")

		case "2":
			fmt.Println("\nYour Tasks:")
			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
			} else {
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}

		case "3":
			fmt.Print("Enter task number to remove: ")
			scanner.Scan()
			index, err := strconv.Atoi(scanner.Text())
			if err != nil || index < 1 || index > len(tasks) {
				fmt.Println("Invalid task number.")
			} else {
				tasks = append(tasks[:index-1], tasks[index:]...)
				fmt.Println("Task removed!")
			}

		case "4":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}
