package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	a.SetIcon(theme.FyneLogo())
	a.Settings().SetTheme(theme.DarkTheme())
	a.Settings().Theme().Font(fyne.TextStyle{Bold: true})
	w := a.NewWindow("RGB to Hex")
	w.SetPadded(true)
	w.Resize(fyne.NewSize(280, 380))

	var filter string = `^([0-9]{1,2}|2[0-4][0-9]|25[0-5])\b+$`

	inputR := &widget.Entry{Validator: validation.NewRegexp(filter, "Must contain a number!")}
	inputR.SetPlaceHolder("Red (0-255)")
	inputG := &widget.Entry{Validator: validation.NewRegexp(filter, "Must contain a number!")}
	inputG.SetPlaceHolder("Green (0-255)")
	inputB := &widget.Entry{Validator: validation.NewRegexp(filter, "Must contain a number!")}
	inputB.SetPlaceHolder("Blue (0-255)")

	title := widget.NewLabel("RGB to Hex")
	footer := widget.NewLabel("Made By Sobek")

	hex := widget.NewEntry()
	hex.SetPlaceHolder("#FFFFFF")
	hex.TextStyle = fyne.TextStyle{Bold: true}

	rgb := container.NewVBox(inputR, inputG, inputB, widget.NewButton("Convert", func() {
		c1, err1 := strconv.ParseInt(inputR.Text, 10, 16)
		if err1 != nil {
		}
		c2, err2 := strconv.ParseInt(inputG.Text, 10, 16)
		if err2 != nil {
		}
		c3, err3 := strconv.ParseInt(inputB.Text, 10, 16)
		if err3 != nil {
		}
		r, g, b := c1, c2, c3
		var str = fmt.Sprintf("#%02X%02X%02X", r, g, b)
		hex.SetText(str)
	}))

	content := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		fyne.NewContainerWithLayout(layout.NewCenterLayout(), title),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), rgb),
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), hex),
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(layout.NewCenterLayout(), footer),
	)

	w.SetContent(content)
	w.CenterOnScreen()
	w.ShowAndRun()
}

/*
Fyne App layout
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	f := app.New()
	w := f.NewWindow("Fyne Testing")
	w.Resize(fyne.NewSize(320, 480))

	title := widget.NewLabel("Title")
	middle := widget.NewLabel("Middle")
	title.Alignment = fyne.TextAlignTrailing
	title.TextStyle = fyne.TextStyle{Italic: true}
	footer := widget.NewLabel("Footer")

	content :=
		fyne.NewContainerWithLayout(
			fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), title, layout.NewSpacer()),
			layout.NewSpacer(),
			fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), middle, layout.NewSpacer()),
			layout.NewSpacer(),
			fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), footer, layout.NewSpacer()),
		)

	w.SetContent(content)
	w.ShowAndRun()
}
*/
