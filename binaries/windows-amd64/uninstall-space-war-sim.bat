@echo off
setlocal EnableExtensions

echo ============================================================
echo   SPACE WAR SIM - Windows Uninstaller
echo ============================================================
echo.

net session >nul 2>&1
if %errorLevel% neq 0 (
    echo ERROR: Please run as Administrator.
    pause
    exit /b 1
)

set "INSTALL_DIR=C:\Program Files\SpaceWarSim"
set "CONFIG_DIR=C:\ProgramData\space-war-sim"

:: Remove binary
if exist "%INSTALL_DIR%\space-war-sim.exe" (
    del /q "%INSTALL_DIR%\space-war-sim.exe"
    echo [OK] Binary removed
)

:: Remove from PATH
set "PATH_KEY=HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment"
for /f "tokens=2*" %%a in ('reg query "%PATH_KEY%" /v Path') do set "CURRENT_PATH=%%b"
echo %CURRENT_PATH% | find /i "%INSTALL_DIR%" >nul 2>&1
if %errorLevel% equ 0 (
    set "NEW_PATH=%CURRENT_PATH:;%INSTALL_DIR%=%"
    reg add "%PATH_KEY%" /v Path /t REG_EXPAND_SZ /d "%NEW_PATH%" /f >nul
    echo [OK] Removed from PATH
)

:: Remove registry entries
reg delete "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\SpaceWarSim" /f >nul 2>&1
echo [OK] Removed from Add/Remove Programs

:: Ask about configs
echo.
set /p "REMOVE_CONFIGS=Remove config files at %CONFIG_DIR%? [y/N] "
if /i "%REMOVE_CONFIGS%"=="y" (
    rmdir /s /q "%CONFIG_DIR%" 2>nul
    echo [OK] Configs removed
) else (
    echo [SKIP] Configs preserved
)

:: Remove install dir if empty
dir /b "%INSTALL_DIR%" 2>nul | findstr . >nul 2>&1 || (
    rmdir /q "%INSTALL_DIR%" 2>nul
    echo [OK] Install directory removed
)

echo.
echo Space War Sim has been uninstalled.
pause