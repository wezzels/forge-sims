@echo off
REM Cyber Red Team Simulator — Windows uninstaller
REM Run as Administrator

NET SESSION >nul 2>&1
if %ERRORLEVEL% neq 0 (
    echo Error: Run as Administrator
    pause
    exit /b 1
)

set INSTDIR=C:\Program Files\CyberRedTeamSim
set CONFDIR=C:\ProgramData\cyber-redteam-sim

echo Uninstalling Cyber Red Team Simulator...

REM Remove from PATH (best effort)
for /f "tokens=2*" %%a in ('reg query "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v Path 2^>nul') do set SYSPATH=%%b
echo %SYSPATH% | find /i "%INSTDIR%" >nul && (
    set "NEWPATH=%SYSPATH:;%INSTDIR%;=;%"
    reg add "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v Path /d "%NEWPATH%" /f >nul
)

REM Remove registry entries
reg delete "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\CyberRedTeamSim" /f 2>nul

REM Remove files
rd /s /q "%INSTDIR%" 2>nul
rd /s /q "%CONFDIR%" 2>nul

echo.
echo Uninstalled. Configs and data directories removed.
pause