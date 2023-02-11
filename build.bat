CHCP 65001

:: 安装包

wails build -u -platform windows/amd64 -webview2 embed -tags "build" -ldflags "-H=windowsgui -s -w"
