package gc

import (
	"gioui.org/io/system"
	"gioui.org/layout"
)

func (g *GcApp) Gui() {
	g.Context = layout.NewContext(g.Window.Queue())
	for e := range g.Window.Events() {
		if e, ok := e.(system.FrameEvent); ok {
			g.Context.Reset(e.Config, e.Size)
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(g.Context,
				layout.Flexed(1, g.Main()),
				layout.Rigid(g.CommandLine()),
				layout.Rigid(g.BottomButtons()),
			)
			e.Frame(g.Context.Ops)
		}
	}
}

func (g *GcApp) CommandLine() func() {
	return func() {
		g.Theme.H6("CommandLine").Layout(g.Context)
	}
}

func (g *GcApp) Main() func() {
	return func() {
		layout.Flex{}.Layout(g.Context,
			layout.Flexed(0.5, g.Folder()),
		)
	}
}
