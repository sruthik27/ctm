package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
)

const (
	daemonEnvVar = "CTM_DAEMON"
)

var (
	pidFile = filepath.Join(os.TempDir(), "ctm.pid")
	logFile = filepath.Join(os.TempDir(), "ctm.log")
)

func main() {
	// Check if we're running as the daemon
	if os.Getenv(daemonEnvVar) == "1" {
		runDaemon()
		return
	}

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

	// Ensure daemon is running
	if err := ensureDaemonRunning(); err != nil {
		fmt.Fprintf(os.Stderr, "Error ensuring daemon: %v\n", err)
		return
	}

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

// ensureDaemonRunning checks if the daemon is running and starts it if necessary
func ensureDaemonRunning() error {
	if pid, err := os.ReadFile(pidFile); err == nil {
		pidInt, _ := strconv.Atoi(string(pid))
		process, err := os.FindProcess(pidInt)
		if err == nil {
			if err = process.Signal(syscall.Signal(0)); err == nil {
				// Daemon is already running
				return nil
			}
		}
	}

	return startDaemon()
}

// startDaemon starts the daemon process
func startDaemon() error {
	executable, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error finding executable: %v", err)
	}

	cmd := exec.Command(executable)
	cmd.Env = append(os.Environ(), "DAEMON=1")
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error starting daemon: %v", err)
	}

	if err := os.WriteFile(pidFile, []byte(strconv.Itoa(cmd.Process.Pid)), 0644); err != nil {
		return fmt.Errorf("error writing PID file: %v", err)
	}

	return nil
}

// runDaemon runs the background notification process
func runDaemon() {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open log file: %v\n", err)
		return
	}
	defer f.Close()

	log.SetOutput(f)

	// Write PID file and setup cleanup
	pid := strconv.Itoa(os.Getpid())
	if err := os.WriteFile(pidFile, []byte(pid), 0644); err != nil {
		log.Printf("Failed to write PID file: %v\n", err)
		return
	}
	defer os.Remove(pidFile)

	// Handle signals for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		log.Printf("Received signal %v, exiting...", sig)
		err := os.Remove(pidFile)
		if err != nil {
			return
		}
		os.Exit(0)
	}()

	// Start notification loop
	startBackgroundNotifier()
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
