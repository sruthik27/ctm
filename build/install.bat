@echo off
SET "CTM_DIR=%ProgramFiles%\ctm"
IF NOT EXIST "%CTM_DIR%" mkdir "%CTM_DIR%"
copy /Y "ctm.exe" "%CTM_DIR%"
setx PATH "%PATH%;%CTM_DIR%"
echo "Installation complete. Please restart your terminal."
