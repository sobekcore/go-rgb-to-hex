package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	a.SetIcon(resourceIconPng)
	a.Settings().SetTheme(theme.DarkTheme())
	a.Settings().Theme().Font(fyne.TextStyle{Bold: true})
	w := a.NewWindow("RGB to Hex")
	w.SetPadded(true)
	w.Resize(fyne.NewSize(250, 280))
	w.SetFixedSize(true)

	var filter string = `^(1?[0-9]{1,2}|2[0-4][0-9]|25[0-5])\b+$`

	inputR := &widget.Entry{Validator: validation.NewRegexp(filter, "Must contain a number!")}
	inputR.SetPlaceHolder("Red (0-255)")
	inputG := &widget.Entry{Validator: validation.NewRegexp(filter, "Must contain a number!")}
	inputG.SetPlaceHolder("Green (0-255)")
	inputB := &widget.Entry{Validator: validation.NewRegexp(filter, "Must contain a number!")}
	inputB.SetPlaceHolder("Blue (0-255)")

	inputHex := &widget.Entry{Validator: validation.NewRegexp(`^#(?:[0-9a-fA-F]{6})$`, "Must be Hex!")}
	inputHex.SetPlaceHolder("#000000")

	title := widget.NewLabel("RGB to Hex")
	titleAlt := widget.NewLabel("Hex to RGB                     ") // White-space (21 spaces) to create title offset
	footer := widget.NewLabel("Made By Sobek")
	footerInfo := widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
		dialog.ShowInformation("About", "RGB to Hex is a simple app\n to convert your given color\n between RGB and Hex. It's core\n is made with Go, and the\n GUI is made with Fyne.io.\n\nAuthor: Sobek", w)
	})

	hexString := widget.NewEntry()
	hexString.TextStyle = fyne.TextStyle{Bold: true}

	rgbRed := widget.NewEntry()
	rgbGreen := widget.NewEntry()
	rgbBlue := widget.NewEntry()
	rgbRed.TextStyle = fyne.TextStyle{Bold: true}
	rgbGreen.TextStyle = fyne.TextStyle{Bold: true}
	rgbBlue.TextStyle = fyne.TextStyle{Bold: true}

	// RGB Fields calculation logic
	rgbFields := container.NewVBox(inputR, inputG, inputB, widget.NewButton("Convert", func() {
		c1, err1 := strconv.ParseInt(inputR.Text, 10, 16)
		if err1 != nil || c1 > 255 || c1 < 0 {
			w.Close()
		}
		c2, err2 := strconv.ParseInt(inputG.Text, 10, 16)
		if err2 != nil || c2 > 255 || c2 < 0 {
			w.Close()
		}
		c3, err3 := strconv.ParseInt(inputB.Text, 10, 16)
		if err3 != nil || c3 > 255 || c3 < 0 {
			w.Close()
		}
		r, g, b := c1, c2, c3
		var str = fmt.Sprintf("#%02X%02X%02X", r, g, b)
		hexString.SetText(str)

		line := &canvas.Rectangle{StrokeColor: color.NRGBA{uint8(r), uint8(g), uint8(b), 255}, StrokeWidth: 4}
		canvas.Refresh(line)

		contentFinal := fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			fyne.NewContainerWithLayout(layout.NewCenterLayout(), title),
			layout.NewSpacer(),
			widget.NewButtonWithIcon("Copy to clipboard", theme.ContentCopyIcon(), func() {
				w.Clipboard().SetContent(str)
				w.Close()
			}),
			fyne.NewContainerWithLayout(layout.NewVBoxLayout(), hexString),
			fyne.NewContainerWithLayout(layout.NewVBoxLayout(), line),
			layout.NewSpacer(),
			fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(),
				footer, layout.NewSpacer(), footerInfo),
		)

		w.SetContent(contentFinal)
	}))

	// Hex Fields calculation logic
	hexFields := container.NewVBox(inputHex, widget.NewButton("Convert", func() {
		r := inputHex.Text[1:3]
		g := inputHex.Text[3:5]
		b := inputHex.Text[5:7]

		c1, err1 := strconv.ParseInt(r, 16, 10)
		if err1 != nil || len(inputHex.Text) != 7 {
			w.Close()
		}
		c2, err2 := strconv.ParseInt(g, 16, 10)
		if err2 != nil || len(inputHex.Text) != 7 {
			w.Close()
		}
		c3, err3 := strconv.ParseInt(b, 16, 10)
		if err3 != nil || len(inputHex.Text) != 7 {
			w.Close()
		}

		var rgb1 = fmt.Sprintf("%v", c1)
		var rgb2 = fmt.Sprintf("%v", c2)
		var rgb3 = fmt.Sprintf("%v", c3)
		rgbRed.SetText("R: " + rgb1)
		rgbGreen.SetText("G: " + rgb2)
		rgbBlue.SetText("B: " + rgb3)

		line := &canvas.Rectangle{StrokeColor: color.NRGBA{uint8(c1), uint8(c2), uint8(c3), 255}, StrokeWidth: 4}
		canvas.Refresh(line)

		contentFinal := fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			fyne.NewContainerWithLayout(layout.NewCenterLayout(), titleAlt),
			layout.NewSpacer(),
			widget.NewButtonWithIcon("Copy to clipboard", theme.ContentCopyIcon(), func() {
				clipboard := fmt.Sprintf("rgb(%v, %v, %v)", rgb1, rgb2, rgb3)
				w.Clipboard().SetContent(clipboard)
				w.Close()
			}),
			fyne.NewContainerWithLayout(layout.NewVBoxLayout(), rgbRed, rgbGreen, rgbBlue),
			fyne.NewContainerWithLayout(layout.NewVBoxLayout(), line),
			layout.NewSpacer(),
			fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(),
				footer, layout.NewSpacer(), footerInfo),
		)

		w.SetContent(contentFinal)
	}))

	// Hex to RGB page render
	contentHex := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		fyne.NewContainerWithLayout(layout.NewCenterLayout(), titleAlt),
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), hexFields),
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(),
			footer, layout.NewSpacer(), footerInfo),
	)

	changeHex := widget.NewButtonWithIcon("Hex", theme.ColorPaletteIcon(), func() {
		w.SetContent(contentHex)
	})

	// RGB to Hex page render
	contentRGB := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(),
			title, layout.NewSpacer(), changeHex),
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), rgbFields),
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(),
			footer, layout.NewSpacer(), footerInfo),
	)

	// Final rendering and app running
	w.SetContent(contentRGB)
	w.CenterOnScreen()
	w.ShowAndRun()
}
