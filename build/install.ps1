# Installation script for Windows
$source = "ctm.exe"
$destination = "$env:ProgramFiles\ctm.exe"

# Check if `ctm.exe` exists
if (!(Test-Path $source)) {
    Write-Output "Error: 'ctm.exe' binary not found. Make sure to build the project first."
    exit 1
}

# Copy to Program Files and set executable permissions
Copy-Item -Path $source -Destination $destination -Force
[System.IO.File]::SetAttributes($destination, 'ReadOnly')

Write-Output "Installation complete. You can now use 'ctm' from anywhere in the terminal."
Write-Output "If the command does not work, ensure the Program Files directory is in your PATH."

# Autocompletion function
$completionScript = @'
# PowerShell autocompletion for ctm
Register-ArgumentCompleter -Native -CommandName ctm -ScriptBlock {
    param($commandName, $wordToComplete, $cursorPosition)
    $commands = @("-a", "-l", "-v", "-r", "-c", "-u", "-p", "-s")
    $commands | Where-Object { $_ -like "$wordToComplete*" } | ForEach-Object {
        [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterName', $_)
    }
}
'@

# Add the autocompletion function to profile
$profilePath = [System.IO.Path]::Combine($HOME, 'Documents\PowerShell\Microsoft.PowerShell_profile.ps1')
if (!(Test-Path -Path $profilePath)) {
    New-Item -ItemType File -Path $profilePath -Force
}

Add-Content -Path $profilePath -Value $completionScript
Write-Output "Autocompletion for 'ctm' added to PowerShell profile."
Write-Output "Please restart PowerShell for autocompletion to take effect."
