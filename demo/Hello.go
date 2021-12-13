// Package main loads a very basic Hello World graphical application.
package main

import (
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/sxn0508/fyne_zh_CN"
)

func main() {
	var lblTxts = []string{
		"Fyne默认是不支持中文的.",
		"可以这样使用中文字体 os.Setenv(\"FYNE_FONT\",\"wqy-microhei.ttf\")",
		"但是你你得先确保系统里有这个字体。",
		"因此我用statik把wqy-microhei字体打包成静态资源。",
		"以后再用只需要引一下包就可以了。像这样：",
		"import _ \"github.com/sxn0508/fyne_zh_CN\"",
	}

	a := app.New()
	w := a.NewWindow("『你好，世界』")
	w.Resize(fyne.Size{400, 100})
	hello := widget.NewLabel("欢迎在Fyne中使用中文!")
	n := 0
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText(lblTxts[n])
			n = (n + 1) % len(lblTxts)
		}),
	))

	w.ShowAndRun()
}
