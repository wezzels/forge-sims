@echo off
setlocal EnableExtensions

echo ============================================================
echo   MARITIME SIM - Windows Installer
echo ============================================================
echo.

:: Check for admin
net session >nul 2>&1
if %errorLevel% neq 0 (
    echo ERROR: Please run as Administrator.
    echo Right-click this file and select "Run as administrator".
    pause
    exit /b 1
)

set "INSTALL_DIR=C:\Program Files\MaritimeSim"
set "CONFIG_DIR=C:\ProgramData\maritime-sim\configs"

:: Create directories
if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"
if not exist "%CONFIG_DIR%" mkdir "%CONFIG_DIR%"

:: Copy binary
copy /Y maritime-sim.exe "%INSTALL_DIR%\maritime-sim.exe" >nul
echo [OK] Binary installed to %INSTALL_DIR%

:: Copy configs
copy /Y configs\*.yaml "%CONFIG_DIR%\" >nul 2>&1
echo [OK] Configs installed to %CONFIG_DIR%

:: Add to PATH
set "PATH_KEY=HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment"
reg query "%PATH_KEY%" /v Path | find /i "%INSTALL_DIR%" >nul 2>&1
if %errorLevel% neq 0 (
    for /f "tokens=2*" %%a in ('reg query "%PATH_KEY%" /v Path') do set "CURRENT_PATH=%%b"
    reg add "%PATH_KEY%" /v Path /t REG_EXPAND_SZ /d "%CURRENT_PATH%;%INSTALL_DIR%" /f >nul
    echo [OK] Added to system PATH
) else (
    echo [SKIP] Already in PATH
)

:: Register in Add/Remove Programs
reg add "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\MaritimeSim" /v DisplayName /d "Maritime Sim" /f >nul
reg add "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\MaritimeSim" /v DisplayVersion /d "0.1.0" /f >nul
reg add "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\MaritimeSim" /v Publisher /d "STSGYM" /f >nul
reg add "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\MaritimeSim" /v UninstallString /d "%INSTALL_DIR%\uninstall.bat" /f >nul
echo [OK] Registered in Add/Remove Programs

echo.
echo ============================================================
echo   Installation complete!
echo ============================================================
echo.
echo   Quick start:
echo     maritime-sim --version
echo     maritime-sim --doctor
echo     maritime-sim --config %CONFIG_DIR%\scenario.yaml
echo.
echo   NOTE: You may need to open a new terminal for PATH to take effect.
echo.
pause