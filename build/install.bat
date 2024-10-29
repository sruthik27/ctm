@echo off
SET "CTM_DIR=%ProgramFiles%\ctm"

:: Create directory if not exists
IF NOT EXIST "%CTM_DIR%" mkdir "%CTM_DIR%"

:: Copy the executable
copy /Y "ctm.exe" "%CTM_DIR%"

:: Add to system PATH (using registry)
reg add "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v Path /t REG_EXPAND_SZ /d "%PATH%;%CTM_DIR%" /f

echo Installation complete. Please restart your terminal.
pause
