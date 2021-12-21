# fyne_zh_CN
fyne中使用中文
Fyne默认是不支持中文的.  
可以这样使用中文字体 
```
os.Setenv(\"FYNE_FONT\",\"wqy-microhei.ttf\")"
```
但是你你得先确保系统里有这个字体。
因此我用statik把wqy-microhei字体打包成静态资源。
以后再用只需要引一下包就可以了。像这样：
```golang
import _ \"github.com/sxn0508/fyne_zh_CN\""
```
完整代码如下
```golang
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/sxn0508/fyne_zh_CN"
)

func main() {
	a := app.New()
	w := a.NewWindow("第一个Fyne程序")

	hello := widget.NewLabel("你好 Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("欢迎使用 :)")
		}),
	))

	w.ShowAndRun()
}
```
