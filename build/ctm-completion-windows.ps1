# Save this as 'ctm-completion.ps1' and dot-source it in your PowerShell profile
Register-ArgumentCompleter -Native -CommandName ctm -ScriptBlock {
    param($wordToComplete, $commandAst, $cursorPosition)

    # Define the available options
    $options = @(
        '-a',
        '-l',
        '-v',
        '-r',
        '-c',
        '-u',
        '-p',
        '-s'
    )

    # Get all task IDs for relevant commands
    $taskIds = if ($commandAst.ToString() -match '-[rcu]') {
        # Execute ctm -l and parse the output for IDs
        $output = & ctm -l
        if ($output) {
            $output | ForEach-Object {
                if ($_ -match '^\d+') {
                    $matches[0]
                }
            }
        }
    }

    # Handle priority values for -p option
    $priorities = if ($commandAst.ToString() -match '-p') {
        1..3  # Returns array of 1,2,3
    }

    # Complete options
    $completions = if ($wordToComplete -match '^-') {
        $options | Where-Object { $_ -like "$wordToComplete*" } | ForEach-Object {
            [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterName', $_)
        }
    }
    # Complete task IDs
    elseif ($taskIds) {
        $taskIds | Where-Object { $_ -like "$wordToComplete*" } | ForEach-Object {
            [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
        }
    }
    # Complete priorities
    elseif ($priorities) {
        $priorities | Where-Object { $_ -like "$wordToComplete*" } | ForEach-Object {
            [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
        }
    }

    # Return all completions
    $completions
}
