# Enabling Auto-Completion for the CTM App

## For macOS/Linux:

1. Open your `~/.zshrc` or `~/.bashrc` file using an editor.
2. Copy all the content from the `ctm-completion-unix.sh` file located in the `build` folder and place it at the end of the `~/.zshrc` or `~/.bashrc` file.
3. Save the file and then run `source ~/.zshrc` or `source ~/.bashrc`.

## For Windows:

1. Copy all the content from the `ctm-completion-windows.ps1` file located in the `build` folder.
2. Add it to your PowerShell profile by following the steps below in the terminal:

```powershell
# First, check if you have a profile
Test-Path $PROFILE

# If it returns false, create a profile
if (!(Test-Path $PROFILE)) {
    New-Item -Type File -Path $PROFILE -Force
}

# Open the profile in notepad to edit it
notepad $PROFILE

# Add the copied content to end of your profile file
