export GO111MODULE=on

-include .env

#编译 Linux 版本
BuildLinux:
	@go build -o micro $(BuildDir) .

#编译 Windows 版本
BuildWindows:
	@go build -o micro.exe $(BuildDir) .

#编译所有环境版本
BuildAll: BuildLinux BuildWindows
