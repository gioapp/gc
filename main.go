package main

import (
	"gioui.org/app"
	"github.com/gioapp/gc/app"
)

func main() {
	go func() {
		g := gc.NewGcApp()
		g.Gui()
	}()
	app.Main()
}
