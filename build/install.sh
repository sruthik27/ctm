#!/bin/bash

# Installation script for macOS/Linux
set -e

# Check if `ctm` binary exists
if [[ ! -f "ctm" ]]; then
  echo "Error: 'ctm' binary not found. Make sure to build the project first."
  exit 1
fi

# Move binary to /usr/local/bin for global access
sudo cp ctm /usr/local/bin/
sudo chmod +x /usr/local/bin/ctm

echo "Installation complete. You can now use 'ctm' from anywhere in the terminal."

#To install the CLI, run the following command:
#$ ./build/install.sh

#To test the CLI, run the following command:
#$ ctm -l

#To add auto-complete of ctm commands append the below to your .~/zshrc or ~/.bashrc file
<< 'SCRIPT_STARTS_FROM_NEXT_LINE'
_ctm_completion() {
    local curcontext="$curcontext" state line
    typeset -A opt_args

    _arguments -C \
        '-a[Add a new task]:task:' \
        '-l[List all tasks]' \
        '-v[View all archived tasks]' \
        '-r[Remove a task by ID]:task ID:->ids' \
        '-c[Complete a task by ID]:task ID:->ids' \
        '-u[Update task priority by ID]:task ID:->ids' \
        '-p[Set or update priority (1=High, 2=Medium, 3=Low)]:priority:(1 2 3)' \
        '-s[Show summary of task completion]'\
        '*::arg:->arg'

    case $state in
        ids)
            # Complete with the task IDs
            ids=($(ctm -l | grep -oP '^\d+')) # Extract task IDs
            _values 'task ID' "${ids[@]}"
            ;;
    esac
}

# Register the completion function for ctm
compdef _ctm_completion ctm
