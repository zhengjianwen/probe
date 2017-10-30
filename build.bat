@echo off
set CGO_ENABLED=0
set GOROOT_BOOTSTRAP=C:/Go
set GOPATH=D:\workspace

echo building...
set GOARCH=amd64
set GOOS=linux
go build
echo building... ok
pause