package main

import (
	"ctm/model"
	"ctm/utils"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func NotifyOverdueTasks() {
	tasks, err := utils.ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	for _, task := range tasks {
		if !task.Done && time.Since(task.CreatedAt) > 24*time.Hour {
			sendNotification(task)
		}
	}
}

func sendNotification(task model.Task) {
	message := fmt.Sprintf("Task '%s' is overdue! Created on %s", task.Name, task.CreatedAt.Format("2006-01-02"))

	switch runtime.GOOS {
	case "darwin":
		cmd := exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "Task Manager"`, message))
		err := cmd.Run()
		if err != nil {
			fmt.Println("Failed to send notification on macOS:", err)
		}
	case "windows":
		cmd := exec.Command("powershell", "-Command", fmt.Sprintf(`[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null; $template = [Windows.UI.Notifications.ToastNotificationManager]::GetTemplateContent([Windows.UI.Notifications.ToastTemplateType]::ToastText02); $textNodes = $template.GetElementsByTagName("text"); $textNodes.Item(0).AppendChild($template.CreateTextNode("Task Manager")); $textNodes.Item(1).AppendChild($template.CreateTextNode("%s")); $toast = [Windows.UI.Notifications.ToastNotification]::new($template); [Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier("Task Manager").Show($toast);`, message))
		err := cmd.Run()
		if err != nil {
			fmt.Println("Failed to send notification on Windows:", err)
		}
	default:
		fmt.Println("Notifications are not supported on this OS.")
	}
}

func startBackgroundNotifier() {
	// Set up signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create ticker for periodic notifications
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	// Detach from parent process
	if runtime.GOOS != "windows" {
		detachProcess()
	}

	for {
		select {
		case <-ticker.C:
			NotifyOverdueTasks()
		case <-sigChan:
			fmt.Println("Daemon shutting down...")
			return
		}
	}
}
