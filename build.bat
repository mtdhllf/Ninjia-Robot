:: �رտ���̨����
@echo off

echo ��������app.json
go build github.com/Tnze/CoolQ-Golang-SDK/tools/cqcfg
go generate
IF ERRORLEVEL 1 pause

echo �������û�������
SET CGO_LDFLAGS=-Wl,--kill-at
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=386
SET GOPROXY=https://goproxy.cn

echo ���ڱ���app.dll
go build -buildmode=c-shared -o app.dll
IF ERRORLEVEL 1 pause

:: ��������˻������������app.dll��app.json���Ƶ���Q��dev�ļ���
REM SET DevDir=D:\��Q Pro\dev\me.cqp.tnze.demo
if defined DevDir (
    echo ���ڸ����ļ�
    for %%f in (app.dll,app.json) do move %%f "%DevDir%\%%f" > nul
    IF ERRORLEVEL 1 pause
)