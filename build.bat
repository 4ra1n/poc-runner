@ECHO OFF

REM SET VERSION
SET "VERSION=0.0.2"
ECHO [*] BUILD POC RUNNER %VERSION%

REM CLEAN BUILD DIR
SET "build_dir=build"
IF EXIST "%build_dir%" (
    IF EXIST "%build_dir%\*" (
        DEL /q /f "%build_dir%\*" 2>nul
    )
    RMDIR /q /s "%build_dir%" 2>nul
)
ECHO [*] CLEAN BUILD DIR FINISH

REM GET NOW TIME
FOR /f "tokens=1-4 delims=/ " %%a IN ('date /t') DO SET MYDATE=%%a/%%b/%%c
FOR /f "tokens=1-4 delims=:., " %%a IN ('time /t') DO SET MYTIME=%%a/%%b
SET NOW=%MYDATE%/%MYTIME%
ECHO [*] NOW TIME %NOW%

REM BUILD GOX
CD gox
go build -o ..\cmd\gox.exe
CD ..\cmd
gox.exe -osarch="darwin/arm64 darwin/amd64 linux/386 linux/amd64 linux/arm linux/arm64 windows/arm windows/arm64 windows/386 windows/amd64" -ldflags="-extldflags=-static -s -w -X main.version=%VERSION% -X main.now=%NOW%" -output="../build/poc-runner-%VERSION%-{{.OS}}-{{.Arch}}"
DEL gox.exe
CD ..
ECHO [*] POC RUNNER BUILD FINISH

REM RUN UPX.EXE - NOT SUPPORT MACOS
bin\upx.exe -qqq build\poc-runner-%VERSION%-darwin-amd64
bin\upx.exe -qqq build\poc-runner-%VERSION%-darwin-arm64
bin\upx.exe -qqq build\poc-runner-%VERSION%-linux-386
bin\upx.exe -qqq build\poc-runner-%VERSION%-linux-amd64
bin\upx.exe -qqq build\poc-runner-%VERSION%-linux-arm
bin\upx.exe -qqq build\poc-runner-%VERSION%-linux-arm64
bin\upx.exe -qqq build\poc-runner-%VERSION%-windows-386.exe
bin\upx.exe -qqq build\poc-runner-%VERSION%-windows-amd64.exe
bin\upx.exe -qqq build\poc-runner-%VERSION%-windows-arm.exe
bin\upx.exe -qqq build\poc-runner-%VERSION%-windows-arm64.exe
bin\upx.exe -qqq build\poc-runner-%VERSION%-freebsd-386
bin\upx.exe -qqq build\poc-runner-%VERSION%-freebsd-amd64
bin\upx.exe -qqq build\poc-runner-%VERSION%-freebsd-arm
bin\upx.exe -qqq build\poc-runner-%VERSION%-openbsd-386
bin\upx.exe -qqq build\poc-runner-%VERSION%-openbsd-amd64
bin\upx.exe -qqq build\poc-runner-%VERSION%-openbsd-arm
bin\upx.exe -qqq build\poc-runner-%VERSION%-solaris-amd64
ECHO [*] POC RUNNER UPX FINISH
