package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var a fyne.App
var w fyne.Window

func main() {
	a = app.New()
	w = a.NewWindow("")
	w.SetTitle("Calculator")
	w.Resize(fyne.NewSize(325, 400))
	design()

	defer w.ShowAndRun()
}

func design() {
	var val1 int = 0
	var val2 int = 0
	var total string = ""
	var symbol string = ""
	var entryText string = ""
	a.Settings().SetTheme(theme.DarkTheme())
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter Values Here ")
	// entry.TextStyle.Italic = true
	entry.TextStyle.Bold = true
	result := canvas.NewText("Result", color.White)
	result.TextStyle.Bold = true
	result.TextSize = 30
	result.Alignment = fyne.TextAlignCenter

	btn7 := widget.NewButton("7", func() {
		entryText = entryText + "7"
		total = total + "7"
		entry.SetText(entryText)
	})
	btn8 := widget.NewButton("8", func() {
		entryText = entryText + "8"
		total = total + "8"
		entry.SetText(entryText)
	})
	btn9 := widget.NewButton("9", func() {
		entryText = entryText + "9"
		total = total + "9"
		entry.SetText(entryText)
	})
	btn_prd := widget.NewButton("*", func() {
		val1, _ = strconv.Atoi(total)
		total = ""
		symbol = "*"
		entryText = entryText + "*"
		entry.SetText(fmt.Sprint(entryText))
	})
	btn4 := widget.NewButton("4", func() {
		entryText = entryText + "4"
		total = total + "4"
		entry.SetText(entryText)
	})
	btn5 := widget.NewButton("5", func() {
		entryText = entryText + "5"
		total = total + "5"
		entry.SetText(entryText)
	})
	btn6 := widget.NewButton("6", func() {
		entryText = entryText + "6"
		total = total + "6"
		entry.SetText(entryText)
	})
	btn_sub := widget.NewButton("-", func() {
		val1, _ = strconv.Atoi(total)
		total = ""
		symbol = "-"
		entryText = entryText + "-"
		entry.SetText(fmt.Sprint(entryText))
	})
	btn1 := widget.NewButton("1", func() {
		entryText = entryText + "1"
		total = total + "1"
		entry.SetText(entryText)
	})
	btn2 := widget.NewButton("2", func() {
		entryText = entryText + "2"
		total = total + "2"
		entry.SetText(entryText)
	})
	btn3 := widget.NewButton("3", func() {
		entryText = entryText + "3"
		total = total + "3"
		entry.SetText(entryText)
	})
	btn_plus := widget.NewButton("+", func() {
		val1, _ = strconv.Atoi(total)
		total = ""
		symbol = "+"
		entryText = entryText + "+"
		entry.SetText(fmt.Sprint(entryText))
	})
	btnBack := widget.NewButton("<--", func() {
		if entry.Text != "" {
			if symbol == "+" || symbol == "-" || symbol == "*" || symbol == "/" {
				entryText = entryText[0 : len(entryText)-1]
				entry.SetText(entryText)
			} else {
				total = total[0 : len(total)-1]
				entryText = entryText[0 : len(entryText)-1]
				entry.SetText(entryText)
			}
		}
	})
	btn0 := widget.NewButton("0", func() {
		entryText = entryText + "0"
		total = total + "0"
		entry.SetText(entryText)
	})
	btndevide := widget.NewButton("/", func() {
		val1, _ = strconv.Atoi(total)
		total = ""
		symbol = "/"
		entryText = entryText + "/"
		entry.SetText(fmt.Sprint(entryText))
	})
	btnEquals := widget.NewButton("=", func() {
		val2, _ = strconv.Atoi(total)
		if symbol == "*" {
			my := val1 * val2
			result.Text = (fmt.Sprint(my))
			// total = ""
			result.Refresh()
		} else if symbol == "+" {
			my := val1 + val2
			result.Text = (fmt.Sprint(my))
			// total = ""
			result.Refresh()
		} else if symbol == "-" {
			my := val1 - val2
			result.Text = (fmt.Sprint(my))
			// total = ""
			result.Refresh()
		} else if symbol == "/" {
			my := val1 / val2
			result.Text = (fmt.Sprint(my))
			// total = ""
			result.Refresh()
		}
	})
	btnClear := widget.NewButton("Clear", func() {
		total = ""
		entryText = ""
		// entry.SetText("")
		entry.SetPlaceHolder("Enter Values Here ")
		result.Text = "Result"
		entry.SetText("")
		result.Refresh()
	})
	// btnClear.FocusGained()
	lay := container.NewGridWithColumns(4,
		btn7,
		btn8,
		btn9,
		btn_prd,
		btn4,
		btn5,
		btn6,
		btn_sub,
		btn1,
		btn2,
		btn3,
		btn_plus,
		btnBack,
		btn0,
		btndevide,
		btnEquals,
	)
	lay.Size().Components()
	vbox := container.NewVBox(result, entry, lay, btnClear)
	w.SetContent(vbox)
}
