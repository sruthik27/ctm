package utils

import (
	"ctm/model"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	taskDirName       = ".ctm"
	taskFileName      = "tasks.json"
	archiveFileName   = "archived.json"
	idCounterFileName = "id_counter.json"
)

func getTaskFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to find home directory: %w", err)
	}
	return filepath.Join(homeDir, taskDirName, taskFileName), nil
}

func getArchiveFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to find home directory: %w", err)
	}
	return filepath.Join(homeDir, taskDirName, archiveFileName), nil
}

func EnsureTaskDir() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("unable to find home directory: %w", err)
	}
	taskDir := filepath.Join(homeDir, taskDirName)
	return os.MkdirAll(taskDir, 0755)
}

func ReadTasks() ([]model.Task, error) {
	if err := EnsureTaskDir(); err != nil {
		return nil, err
	}

	taskFilePath, err := getTaskFilePath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(taskFilePath); os.IsNotExist(err) {
		return []model.Task{}, nil
	}

	file, err := os.ReadFile(taskFilePath)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	if err := json.Unmarshal(file, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func WriteTasks(tasks []model.Task) error {
	if err := EnsureTaskDir(); err != nil {
		return err
	}

	taskFilePath, err := getTaskFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return os.WriteFile(taskFilePath, data, 0644)
}

func ReadArchivedTasks() ([]model.Task, error) {
	archiveFilePath, err := getArchiveFilePath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(archiveFilePath); os.IsNotExist(err) {
		return []model.Task{}, nil
	}

	file, err := os.ReadFile(archiveFilePath)
	if err != nil {
		return nil, err
	}

	var archivedTasks []model.Task
	if err := json.Unmarshal(file, &archivedTasks); err != nil {
		return nil, err
	}
	return archivedTasks, nil
}

func WriteArchivedTasks(tasks []model.Task) error {
	archiveFilePath, err := getArchiveFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return os.WriteFile(archiveFilePath, data, 0644)
}

// UpdateTaskIDCounter Helper function to update the ID counter in the id_counter.json file
func UpdateTaskIDCounter(idCounter int) error {
	counterFilePath := GetIDCounterFilePath()
	data, err := json.Marshal(idCounter)
	if err != nil {
		return err
	}
	return os.WriteFile(counterFilePath, data, 0644)
}

// FindTaskIndexByID Helper function to find the index of a task by its ID
func FindTaskIndexByID(tasks []model.Task, id int) (int, error) {
	for i, task := range tasks {
		if task.ID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("task with ID %d not found", id)
}

// GetIDCounterFilePath Helper function to get the path to the id_counter.json file
func GetIDCounterFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: unable to find home directory.")
		os.Exit(1)
	}
	return filepath.Join(homeDir, taskDirName, idCounterFileName)
}
