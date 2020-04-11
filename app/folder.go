package gc

import (
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"time"
)

func (g *GcApp) Folder() func() {
	return func() {
		layout.Flex{
			Axis: layout.Vertical,
		}.Layout(g.Context,
			layout.Rigid(g.Pwd()),
			layout.Flexed(1, g.FilesList()),
			layout.Rigid(g.CursorPosition()),
		)
	}
}

func (g *GcApp) Pwd() func() {
	return func() {
		g.Theme.H6("C:\\FolderPath").Layout(g.Context)
	}
}

func (g *GcApp) CursorPosition() func() {
	return func() {
		g.Theme.H6(".. UP-DIR 6-6-98").Layout(g.Context)
	}
}

func (g *GcApp) FilesList() func() {
	return func() {
		listThings("/", g.Context)
	}
}

type Thing struct {
	Name     string
	Type     string
	Size     int64
	ModTime  time.Time
	out      interface{}
	pressed  bool
	selected bool
}

func listThings(dir string, gtx *layout.Context) {
	list := &layout.List{
		Axis: layout.Vertical,
	}
	if dir == "" {
		dir = "/"
	}
	listThings, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	things := make(map[string]*Thing)
	for _, t := range listThings {
		things[t.Name()] = &Thing{
			Name: t.Name(),
		}
		if t.Mode().IsRegular() {
			things[t.Name()].Size = t.Size()
			things[t.Name()].ModTime = t.ModTime()
		}
		if t.IsDir() {
			things[t.Name()].Type = "folder"
		}
	}

	list.Layout(gtx, len(listThings), func(i int) {
		col := color.RGBA{A: 0xff, R: 0xcf, G: 0xcf, B: 0x30}
		if listThings[i].IsDir() {
		} else {
			col = color.RGBA{A: 0xff, R: 0xcf, G: 0x30, B: 0x30}
		}
		things[listThings[i].Name()].Layout(listThings[i].Name(), col, gtx)
	})
}

func (t *Thing) Layout(text string, col color.RGBA, gtx *layout.Context) {
	for _, e := range gtx.Events(t) {
		if e, ok := e.(pointer.Event); ok {
			switch e.Type {
			case pointer.Press:
				t.pressed = true
				t.selected = true
			case pointer.Release:
				t.pressed = false
			}
		}
	}
	th := material.NewTheme()

	if t.pressed {
		col = color.RGBA{A: 0xff, R: 0x30, G: 0xcf, B: 0xcf}
	}
	pointer.Rect(
		image.Rectangle{Max: image.Point{X: 500, Y: 500}},
	).Add(gtx.Ops)
	pointer.InputOp{Key: t}.Add(gtx.Ops)
	drawSquare(gtx.Ops, col)
	th.H6(text).Layout(gtx)
}

func drawSquare(ops *op.Ops, color color.RGBA) {
	square := f32.Rectangle{
		Max: f32.Point{X: 500, Y: 500},
	}
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{Rect: square}.Add(ops)
}
