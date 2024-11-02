package main

import (
	"ctm/model"
	"ctm/utils"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/mattn/go-colorable"
	"os"
)

func ListTasks() {
	tasks, err := utils.ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks available!")
		return
	}

	var highPriorityTasks []model.Task
	var mediumPriorityTasks []model.Task
	var lowPriorityTasks []model.Task

	for _, task := range tasks {
		switch task.Priority {
		case 1:
			highPriorityTasks = append(highPriorityTasks, task)
		case 2:
			mediumPriorityTasks = append(mediumPriorityTasks, task)
		case 3:
			lowPriorityTasks = append(lowPriorityTasks, task)
		}
	}

	displayGroupedTasks("High Priority", highPriorityTasks)
	displayGroupedTasks("Medium Priority", mediumPriorityTasks)
	displayGroupedTasks("Low Priority", lowPriorityTasks)
}

func displayGroupedTasks(title string, tasks []model.Task) {
	if len(tasks) > 0 {
		fmt.Printf("\n%s:\n", title)

		var incompleteTasks []model.Task
		var completedTasks []model.Task
		for _, task := range tasks {
			if task.Done {
				completedTasks = append(completedTasks, task)
			} else {
				incompleteTasks = append(incompleteTasks, task)
			}
		}

		for _, task := range incompleteTasks {
			displayTask(task)
		}

		for _, task := range completedTasks {
			displayTask(task)
		}
	}
}

var out = colorable.NewColorableStdout()

func displayTask(task model.Task) {
	status := "\033[0m" + task.Name
	if task.Done {
		status = "\033[9m" + task.Name + "\033[0m"
	}

	priorityStar := getPriorityStar(task.Priority)
	fmt.Fprintf(out, "%s %s [%d]\n", priorityStar, status, task.ID)
}

func getPriorityStar(priority int) string {
	switch priority {
	case 1:
		return "\033[31m★\033[0m"
	case 2:
		return "\033[33m★\033[0m"
	case 3:
		return "\033[32m★\033[0m"
	}
	return ""
}

func ListTasksInTableView() {
	tasks, err := utils.ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks available!")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.AppendHeader(table.Row{"ID", "Priority", "Status", "Task Name", "Created At"})

	var highPriorityTasks []model.Task
	var mediumPriorityTasks []model.Task
	var lowPriorityTasks []model.Task

	for _, task := range tasks {
		switch task.Priority {
		case 1:
			highPriorityTasks = append(highPriorityTasks, task)
		case 2:
			mediumPriorityTasks = append(mediumPriorityTasks, task)
		case 3:
			lowPriorityTasks = append(lowPriorityTasks, task)
		}
	}

	addTasksToTable := func(tasks []model.Task) {
		for _, task := range tasks {
			if !task.Done {
				t.AppendRow(table.Row{
					task.ID,
					getPriorityDisplay(task.Priority),
					text.Colors{text.FgYellow}.Sprint("Pending"),
					task.Name,
					task.CreatedAt.Format("2006-01-02"),
				})
			}
		}
		for _, task := range tasks {
			if task.Done {
				t.AppendRow(table.Row{
					task.ID,
					getPriorityDisplay(task.Priority),
					text.Colors{text.FgGreen}.Sprint("Completed"),
					task.Name,
					task.CreatedAt.Format("2006-01-02"),
				})
			}
		}
	}

	addTasksToTable(highPriorityTasks)
	addTasksToTable(mediumPriorityTasks)
	addTasksToTable(lowPriorityTasks)

	t.SetColumnConfigs([]table.ColumnConfig{
		{Name: "ID", AutoMerge: true},
		{Name: "Task Name", WidthMax: 50},
	})

	t.Style().Options.SeparateRows = true
	t.Style().Box.PaddingLeft = "  "
	t.Style().Box.PaddingRight = "  "
	t.Style().Color.Header = text.Colors{text.FgBlue, text.Bold}
	t.Style().Color.Border = text.Colors{text.FgBlue}
	t.Style().Color.Separator = text.Colors{text.FgBlue}
	t.Render()
}

func DisplaySummary() {
	tasks, err := utils.ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	totalTasks := len(tasks)
	if totalTasks == 0 {
		fmt.Println("No tasks available to summarize!")
		return
	}

	completedTasks := 0
	incompleteTasks := 0

	for _, task := range tasks {
		if task.Done {
			completedTasks++
		} else {
			incompleteTasks++
		}
	}

	completedPercentage := (float64(completedTasks) / float64(totalTasks)) * 100
	incompletePercentage := (float64(incompleteTasks) / float64(totalTasks)) * 100

	fmt.Println("Task Summary:")
	fmt.Printf("Total tasks: %d\n", totalTasks)
	fmt.Printf("Completed tasks: %d (%.2f%%)\n", completedTasks, completedPercentage)
	fmt.Printf("Incomplete tasks: %d (%.2f%%)\n", incompleteTasks, incompletePercentage)
}

func ViewArchivedTasks() {
	archivedTasks, err := utils.ReadArchivedTasks()
	if err != nil {
		fmt.Println("Error reading archived tasks:", err)
		return
	}

	if len(archivedTasks) == 0 {
		fmt.Println("No archived tasks available!")
		return
	}

	fmt.Println("Archived Tasks:")
	for _, task := range archivedTasks {
		fmt.Printf("%s - %s\n", task.Name, task.CreatedAt.Format("2006-01-02"))
	}
}

func getPriorityDisplay(priority int) string {
	switch priority {
	case 1:
		return text.Colors{text.FgRed}.Sprint("★")
	case 2:
		return text.Colors{text.FgYellow}.Sprint("★")
	default:
		return text.Colors{text.FgGreen}.Sprint("★")
	}
}
