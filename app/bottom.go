package gc

import (
	"gioui.org/layout"
	"github.com/gioapp/gel"
)

var (
	helpButton  = new(gel.Button)
	menuButton  = new(gel.Button)
	viewButton  = new(gel.Button)
	editButton  = new(gel.Button)
	copyButton  = new(gel.Button)
	moveButton  = new(gel.Button)
	mkdirButton = new(gel.Button)
	delButton   = new(gel.Button)
	quitButton  = new(gel.Button)
	pullButton  = new(gel.Button)
)

func (g *GcApp) BottomButtons() func() {
	return func() {
		layout.Flex{}.Layout(g.Context,
			layout.Flexed(0.1, func() {
				g.Theme.DuoUIbutton(g.Theme.Fonts["Mono"], "Help", g.Theme.Colors["blue-lite-blue"], g.Theme.Colors["Primary"], "", "", "", "", 16, 0, 0, 0, 13, 16, 14, 16).Layout(g.Context, helpButton)
			}),
			layout.Flexed(0.1, func() {
				g.Theme.DuoUIbutton(g.Theme.Fonts["Mono"], "Menu", g.Theme.Colors["blue-lite-blue"], g.Theme.Colors["Primary"], "", "", "", "", 16, 0, 0, 0, 13, 16, 14, 16).Layout(g.Context, menuButton)
			}),
			layout.Flexed(0.1, func() {
				g.Theme.DuoUIbutton(g.Theme.Fonts["Mono"], "View", g.Theme.Colors["blue-lite-blue"], g.Theme.Colors["Primary"], "", "", "", "", 16, 0, 0, 0, 13, 16, 14, 16).Layout(g.Context, viewButton)
			}),
			layout.Flexed(0.1, func() {
				g.Theme.DuoUIbutton(g.Theme.Fonts["Mono"], "Edit", g.Theme.Colors["blue-lite-blue"], g.Theme.Colors["Primary"], "", "", "", "", 16, 0, 0, 0, 13, 16, 14, 16).Layout(g.Context, editButton)
			}),
			layout.Flexed(0.1, func() {
				g.Theme.DuoUIbutton(g.Theme.Fonts["Mono"], "Copy", g.Theme.Colors["blue-lite-blue"], g.Theme.Colors["Primary"], "", "", "", "", 16, 0, 0, 0, 13, 16, 14, 16).Layout(g.Context, copyButton)
			}),
			layout.Flexed(0.1, func() {
				g.Theme.DuoUIbutton(g.Theme.Fonts["Mono"], "Move", g.Theme.Colors["blue-lite-blue"], g.Theme.Colors["Primary"], "", "", "", "", 16, 0, 0, 0, 13, 16, 14, 16).Layout(g.Context, moveButton)
			}),
			layout.Flexed(0.1, func() {
				g.Theme.DuoUIbutton(g.Theme.Fonts["Mono"], "Mkdir", g.Theme.Colors["blue-lite-blue"], g.Theme.Colors["Primary"], "", "", "", "", 16, 0, 0, 0, 13, 16, 14, 16).Layout(g.Context, mkdirButton)
			}),
			layout.Flexed(0.1, func() {
				g.Theme.DuoUIbutton(g.Theme.Fonts["Mono"], "Del", g.Theme.Colors["blue-lite-blue"], g.Theme.Colors["Primary"], "", "", "", "", 16, 0, 0, 0, 13, 16, 14, 16).Layout(g.Context, delButton)
			}),
			layout.Flexed(0.1, func() {
				g.Theme.DuoUIbutton(g.Theme.Fonts["Mono"], "Quit", g.Theme.Colors["blue-lite-blue"], g.Theme.Colors["Primary"], "", "", "", "", 16, 0, 0, 0, 13, 16, 14, 16).Layout(g.Context, quitButton)
			}),
			layout.Flexed(0.1, func() {
				g.Theme.DuoUIbutton(g.Theme.Fonts["Mono"], "Pull", g.Theme.Colors["blue-lite-blue"], g.Theme.Colors["Primary"], "", "", "", "", 16, 0, 0, 0, 13, 16, 14, 16).Layout(g.Context, pullButton)
			}),
		)
	}
}
