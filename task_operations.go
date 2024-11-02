package main

import (
	"ctm/model"
	"ctm/utils"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func AddTask(name string, priority int) {
	tasks, err := utils.ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	if priority < 1 || priority > 3 {
		priority = 3
	}

	newID, err := getNextTaskID()
	if err != nil {
		fmt.Println("Error generating task ID:", err)
		return
	}

	newTask := model.Task{
		ID:        newID,
		Name:      name,
		Done:      false,
		Priority:  priority,
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)
	if err := utils.WriteTasks(tasks); err != nil {
		fmt.Println("Error saving task:", err)
	} else {
		fmt.Printf("Added task: %s with priority %d and ID %d\n", name, priority, newID)
	}
}

func RemoveTask(id int) {
	tasks, err := utils.ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	index, err := utils.FindTaskIndexByID(tasks, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	if err := utils.WriteTasks(tasks); err != nil {
		fmt.Println("Error removing task:", err)
	} else {
		fmt.Printf("Removed task with ID: %d\n", id)
	}
}

func CompleteTask(id int) {
	tasks, err := utils.ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	index, err := utils.FindTaskIndexByID(tasks, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	tasks[index].Done = !tasks[index].Done
	if err := utils.WriteTasks(tasks); err != nil {
		fmt.Println("Error completing task:", err)
	} else {
		if tasks[index].Done {
			fmt.Printf("Completed task with ID: %d\n", id)
		} else {
			fmt.Printf("Task ID %d back on tracking\n", id)
		}
	}
}

func UpdateTaskPriority(id int, priority int) {
	tasks, err := utils.ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	if priority < 1 || priority > 3 {
		fmt.Println("Invalid priority value!")
		return
	}

	index, err := utils.FindTaskIndexByID(tasks, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	tasks[index].Priority = priority
	if err := utils.WriteTasks(tasks); err != nil {
		fmt.Println("Error updating task priority:", err)
	} else {
		fmt.Printf("Updated priority for task with ID: %d to %d\n", id, priority)
	}
}

func archiveOldTasks() {
	tasks, err := utils.ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	var recentTasks []model.Task
	var tasksToArchive []model.Task

	for _, task := range tasks {
		if time.Since(task.CreatedAt).Hours() > 24 && task.Done {
			tasksToArchive = append(tasksToArchive, task)
		} else {
			recentTasks = append(recentTasks, task)
		}
	}

	if len(tasksToArchive) > 0 {
		archivedTasks, err := utils.ReadArchivedTasks()
		if err == nil {
			tasksToArchive = append(archivedTasks, tasksToArchive...)
		}

		if err := utils.WriteArchivedTasks(tasksToArchive); err != nil {
			fmt.Println("Error archiving tasks:", err)
		} else {
			fmt.Printf("Archived %d completed tasks\n", len(tasksToArchive))
		}
	}

	if err := utils.WriteTasks(recentTasks); err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

// Helper function to get the next unique task ID
func getNextTaskID() (int, error) {
	// Ensure the .ctm directory exists
	if err := utils.EnsureTaskDir(); err != nil {
		return 0, err
	}

	// Read the current counter
	counterFilePath := utils.GetIDCounterFilePath()
	var idCounter int
	if _, err := os.Stat(counterFilePath); os.IsNotExist(err) {
		// If the counter file doesn't exist, start from 1
		idCounter = 1
	} else {
		// Read the existing counter value
		data, err := os.ReadFile(counterFilePath)
		if err != nil {
			return 0, err
		}
		if err := json.Unmarshal(data, &idCounter); err != nil {
			return 0, err
		}
		// Increment the counter for the new ID
		idCounter++
	}

	// Update the counter in the file
	if err := utils.UpdateTaskIDCounter(idCounter); err != nil {
		return 0, err
	}

	return idCounter, nil
}
