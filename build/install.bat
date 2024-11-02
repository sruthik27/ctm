@echo off
SET "CTM_DIR=%ProgramFiles%\ctm"

:: Create directory if it does not exist
IF NOT EXIST "%CTM_DIR%" mkdir "%CTM_DIR%"

:: Copy the executable
copy /Y "ctm.exe" "%CTM_DIR%"

:: Add to system PATH (using registry for all users)
reg add "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v Path /t REG_EXPAND_SZ /d "%PATH%;%CTM_DIR%" /f

:: Add to user PATH (for current user only)
FOR /F "usebackq tokens=2* delims= " %%A IN (`reg query "HKCU\Environment" /v Path 2^>nul`) DO SET "USER_PATH=%%B"
IF DEFINED USER_PATH (
    SET "NEW_USER_PATH=%USER_PATH%;%CTM_DIR%"
) ELSE (
    SET "NEW_USER_PATH=%CTM_DIR%"
)

reg add "HKCU\Environment" /v Path /t REG_EXPAND_SZ /d "%NEW_USER_PATH%" /f

echo Installation complete. Please restart your terminal.
pause
