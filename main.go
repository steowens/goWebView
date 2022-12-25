package main

import (
	"log"

	"github.com/webview/webview"
	"golang.org/x/sys/windows"
)

var hModule windows.Handle

const (
	WEBVIEW_HINT_NONE  webview.Hint = 0
	WEBVIEW_HINT_MIN   webview.Hint = 1
	WEBVIEW_HINT_MAX   webview.Hint = 2
	WEBVIEW_HINT_FIXED webview.Hint = 3
)

func init() {
	err := windows.GetModuleHandleEx(0x01, nil, &hModule)
	if err != nil {
		panic(err)
	}

}

func main() {
	log.Println("Hello")
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(480, 320, WEBVIEW_HINT_NONE)
	w.SetHtml("Thanks for using webview!")
	w.Run()
}
