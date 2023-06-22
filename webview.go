//go:build !windows
// +build !windows

package wasabi

import (
	"fmt"

	"github.com/webview/webview"
)

type webView struct {
	wv webview.WebView
}

func (wv *webView) open(title string, port, width, height int) {
	wv.wv.SetTitle(title)
	wv.wv.SetSize(width, height, webview.HintNone)
	wv.wv.Navigate(fmt.Sprintf("http://localhost:%d", port))
	defer wv.wv.Destroy()
	wv.wv.Run()
	return
}

func detectWebview() *webView {
	wv := webview.New(false)
	if wv == nil {
		return nil
	}
	return &webView{wv: wv}
}
