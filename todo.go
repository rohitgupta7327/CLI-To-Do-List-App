package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

var tasks []Task
var nextID = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n1. Add Task\n2. List Tasks\n3. Complete Task\n4. Exit")
		fmt.Print("Choose an option: ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			addTask(scanner)
		case "2":
			listTasks()
		case "3":
			completeTask(scanner)
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task title: ")
	scanner.Scan()
	title := scanner.Text()
	task := Task{ID: nextID, Title: title, Completed: false}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Task added successfully!")
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("\nTo-Do List:")
	for _, task := range tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Title)
	}
}

func completeTask(scanner *bufio.Scanner) {
	listTasks()
	fmt.Print("Enter task ID to mark as completed: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			fmt.Println("Task marked as completed!")
			return
		}
	}
	fmt.Println("Task not found.")
}
