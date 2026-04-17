@echo off
REM Cyber Red Team Simulator — Windows installer
REM Run as Administrator

echo ===============================================
echo   CYBER RED TEAM SIM - Windows Installer
echo ===============================================
echo.

NET SESSION >nul 2>&1
if %ERRORLEVEL% neq 0 (
    echo Error: Run as Administrator
    echo Right-click this file and select "Run as administrator"
    pause
    exit /b 1
)

set BINARY=cyber-redteam-sim.exe
set INSTDIR=C:\Program Files\CyberRedTeamSim
set CONFDIR=C:\ProgramData\cyber-redteam-sim\configs

echo Installing to %INSTDIR%...

REM Create directories
if not exist "%INSTDIR%" mkdir "%INSTDIR%"
if not exist "%CONFDIR%" mkdir "%CONFDIR%"

REM Copy binary
if exist "%~dp0% BINARY%" (
    copy /Y "%~dp0% BINARY%" "%INSTDIR%\%BINARY%" >nul
) else if exist "%~dp0cyber-redteam-sim.exe" (
    copy /Y "%~dp0cyber-redteam-sim.exe" "%INSTDIR%\%BINARY%" >nul
) else (
    echo Error: Cannot find %BINARY% in %~dp0
    pause
    exit /b 1
)

REM Copy configs
if exist "%~dp0configs" (
    copy /Y "%~dp0configs\*.yaml" "%CONFDIR%\" >nul 2>&1
) else if exist "%~dp0etc\cyber-redteam-sim\configs" (
    copy /Y "%~dp0etc\cyber-redteam-sim\configs\*.yaml" "%CONFDIR%\" >nul 2>&1
)

REM Add to PATH
setx PATH "%PATH%;%INSTDIR%" /M >nul 2>&1
echo Added %INSTDIR% to system PATH

REM Register uninstall info
reg add "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\CyberRedTeamSim" /v "DisplayName" /d "Cyber Red Team Simulator" /f >nul
reg add "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\CyberRedTeamSim" /v "DisplayVersion" /d "0.2.0" /f >nul
reg add "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\CyberRedTeamSim" /v "Publisher" /d "STSGYM" /f >nul
reg add "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\CyberRedTeamSim" /v "InstallLocation" /d "%INSTDIR%" /f >nul
reg add "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\CyberRedTeamSim" /v "UninstallString" /d "%INSTDIR%\uninstall.bat" /f >nul

REM Create uninstaller
echo @echo off > "%INSTDIR%\uninstall.bat"
echo echo Uninstalling Cyber Red Team Simulator... >> "%INSTDIR%\uninstall.bat"
echo reg delete "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\CyberRedTeamSim" /f >> "%INSTDIR%\uninstall.bat"
echo rd /s /q "%INSTDIR%" >> "%INSTDIR%\uninstall.bat"
echo rd /s /q "%CONFDIR%\.." >> "%INSTDIR%\uninstall.bat"
echo echo Done. >> "%INSTDIR%\uninstall.bat"

echo.
echo ===============================================
echo   Installation complete!
echo ===============================================
echo.
echo Binary:     %INSTDIR%\%BINARY%
echo Configs:    %CONFDIR%\
echo.
echo Open a NEW command prompt and run:
echo   cyber-redteam-sim --version
echo   cyber-redteam-sim --config "%CONFDIR%\enterprise-ad.yaml" --no-server
echo.
"%INSTDIR%\%BINARY%" --version
pause