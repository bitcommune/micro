AppName := micro
BuildDir := $(CURDIR)/build
ApplicationPath := $(BuildDir)/bin/$(AppName)

# 编译windows
BuildWindows:
	set GOOS=windows
	set GOARCH=amd64
	go build -o $(GOPATH)/bin/$(AppName).exe .


.PHONY: BuildWindows