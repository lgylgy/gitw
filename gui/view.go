package gui

import "github.com/jroimartin/gocui"

type View struct {
	name string
	x0   int
	y0   int
	x1   int
	y1   int
}

func (v *View) GetName() string {
	return v.name
}

func (v *View) Draw(g *gocui.Gui) (*gocui.View, error) {
	return g.SetView(v.name, v.x0, v.y0, v.x1, v.y1)
}
