# CTM - Command Line Task Manager

**CTM** (Command Line Task Manager) is a lightweight CLI-based tool for managing tasks, setting priorities, and organizing to-dos efficiently from the terminal. Designed to be simple yet effective, CTM provides a set of powerful commands for adding, completing, prioritizing, listing, and archiving tasks right from the command line.

## Features

- **Add Tasks**: Quickly add new tasks with options to set priorities.
- **List Tasks**: View active and completed tasks in grouped or tabular views.
- **Prioritize Tasks**: Assign priority levels (1=High, 2=Medium, 3=Low) to tasks.
- **Remove Tasks**: Delete tasks by specifying their IDs.
- **Complete Tasks**: Mark tasks as completed by their IDs.
- **Update Task Priority**: Change the priority level of a task.
- **Task Completion Summary**: Display a summary of tasks being tracked.
- **Archive Completed Tasks**: Automatically archives tasks upon completion after a day.
- **View Archived Tasks**: List all archived tasks for reference.
- **Notifications**: Background reminders for tasks not completed over two days.

## Installation

CTM is available for macOS, Linux, and Windows. Follow the installation instructions below for your specific OS.

### macOS / Linux

1. **Clone the repository**:
   * `git clone <repository_url>`
   * `cd ctm/build`
   
2. **Run the installation script**:
   * `chmod +x build/install.sh`
   * `./install.sh`

   The script will:
   - Move the `ctm` binary to `/usr/local/bin` for global access.

3. **Enable auto completions**:
   Please refer enabling_auto_completion_guide.md file for steps, After installation, reload your terminal or run `source ~/.zshrc` / `source ~/.bashrc` to enable autocompletion.

### Windows

1. **Clone the repository**:
   Open Powershell as administrator
   * `git clone <repository_url>`
   * `cd ctm/build`
   
3. **Run the installation script**:
   `.\install.bat`

   The script will:
    - Move `ctm.exe` to `Program Files` for global access.
   
4. **Enable auto completions**:
   Please refer enabling_auto_completion_guide.md file for steps, After installation, restart your terminal to enable autocompletion.

### Custom Build
If you want to modify the code and deploy your own version, make changes and do `go build -o ctm` (UNIX) 
or `GOOS=windows GOARCH=amd64 go build -o build/ctm.exe` (WINDOWS) and then proceed with installation scripts

## Usage

Below is a summary of the main commands available with CTM.

### Add a New Task

```bash
ctm -a "Your task description here"
```

### To Add a New Task with Priority (default is low)
```bash
ctm -a "Your task description here" -p "Your priority level"
```
Priority should be integer 1 or 2 or 3
- `1` = High
- `2` = Medium
- `3` = Low

### List All Tasks 
#### Lists in grouped list format
```bash
ctm -l
```
#### To list in tabular format also include -t flag
```bash
ctm -l -t
```

### View Archived Tasks

```bash
ctm -v
```

### Remove a Task by ID

```bash
ctm -r <task ID>
```

### Complete a Task by ID

```bash
ctm -c <task ID>
```

### Update Task Priority by ID

```bash
ctm -u <task ID> -p <priority level>
```

- Priority levels:
    - `1` = High
    - `2` = Medium
    - `3` = Low

### Show Task Completion Summary

```bash
ctm -s
```

## Contributing

Contributions are welcome! Feel free to fork this repository and submit a pull request with improvements, optimizations, or new features.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
