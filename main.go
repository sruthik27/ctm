package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	archiveOldTasks()
	// Define and parse CLI flags
	addFlag := flag.String("a", "", "Add a task")
	listFlag := flag.Bool("l", false, "List all tasks")
	tableFlag := flag.Bool("t", false, "Display tasks in table view")
	removeFlag := flag.Int("r", -1, "Remove task by ID")
	completeFlag := flag.Int("c", -1, "Complete task by ID")
	updatePriorityFlag := flag.Int("u", -1, "Update task priority by ID")
	priorityFlag := flag.Int("p", 0, "Set priority (1=High, 2=Medium, 3=Low)")
	viewArchivedFlag := flag.Bool("v", false, "View archived tasks")
	summaryFlag := flag.Bool("s", false, "Display summary of task completion")
	flag.Parse()

	// Validate flag combinations
	if err := validateFlags(*addFlag, *removeFlag, *completeFlag, *updatePriorityFlag, *priorityFlag); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		flag.Usage()
		return
	}

	// Handle commands
	switch {
	case *addFlag != "":
		AddTask(*addFlag, *priorityFlag)
	case *listFlag:
		if *tableFlag {
			ListTasksInTableView()
		} else {
			ListTasks()
		}
	case *removeFlag >= 0:
		RemoveTask(*removeFlag)
	case *completeFlag >= 0:
		CompleteTask(*completeFlag)
	case *updatePriorityFlag >= 0 && *priorityFlag > 0:
		UpdateTaskPriority(*updatePriorityFlag, *priorityFlag)
	case *viewArchivedFlag:
		ViewArchivedTasks()
	case *summaryFlag:
		DisplaySummary()
	default:
		fmt.Println("Invalid command! Use -a, -l, -t, -r, -c, -u, -v, or -s.")
	}
}

// validateFlags checks for incompatible or invalid flag combinations
func validateFlags(add string, remove, complete, updatePriority, priority int) error {
	if add != "" && (remove >= 0 || complete >= 0 || updatePriority >= 0) {
		return errors.New("cannot add a task while removing, completing, or updating another task")
	}
	if remove >= 0 && (add != "" || complete >= 0 || updatePriority >= 0) {
		return errors.New("cannot remove a task while adding, completing, or updating another task")
	}
	if complete >= 0 && (add != "" || remove >= 0 || updatePriority >= 0) {
		return errors.New("cannot complete a task while adding, removing, or updating another task")
	}
	if updatePriority >= 0 && priority == 0 {
		return errors.New("priority must be set when updating task priority")
	}
	return nil
}
