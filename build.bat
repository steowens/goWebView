
set CGO_CXXFLAGS="-I%cd%\webview2\build\native\include"
set CGO_LDFLAGS="-L%cd%\webview2\build\native\x64"

go build  -ldflags="-H windowsgui" -o build/goWebview.exe  

