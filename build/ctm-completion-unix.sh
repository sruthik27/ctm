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
