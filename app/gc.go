package gc

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gelook"
)

type GcApp struct {
	Window  *app.Window
	Context *layout.Context
	Theme   *gelook.DuoUItheme
}

func NewGcApp() *GcApp {
	gofont.Register()
	w := app.NewWindow(
		app.Size(unit.Dp(1000), unit.Dp(800)),
		app.Title("ParallelCoin"),
	)
	return &GcApp{
		Window: w,
		Theme:  gelook.NewDuoUItheme(),
	}
}
