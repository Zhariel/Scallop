package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createWindow(w, h int, name string) fyne.Window {
	a := app.New()
	win := a.NewWindow(name)
	win.Resize(fyne.NewSize(float32(w), float32(h)))
	text := widget.NewLabel("New Content")
	win.SetContent(text)
	return win
}
