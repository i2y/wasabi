//go:build windows
// +build windows

package wasabi

import (
	"fmt"
	"syscall"

	"github.com/jchv/go-webview2"
)

type webView struct {
	wv webview2.WebView
}

func (wv *webView) open(title string, port, width, height int) {
	wv.wv.SetTitle(title)
	wv.wv.SetSize(width, height, webview2.HintNone)
	wv.wv.Navigate(fmt.Sprintf("http://localhost:%d", port))
	defer wv.wv.Destroy()
	wv.wv.Run()
}

func detectWebview() *webView {
	dll := syscall.MustLoadDLL("user32")
	if proc, err := dll.FindProc("SetProcessDpiAwarenessContext"); err == nil {
		aware := -4
		proc.Call(uintptr(aware))
	}
	wv := webview2.New(false)
	if wv == nil {
		return nil
	}
	return &webView{wv: wv}
}
