package gui

import "github.com/jroimartin/gocui"

type View struct {
	name string
	x0   float32
	y0   float32
	x1   float32
	y1   float32
}

func (v *View) GetName() string {
	return v.name
}

func (v *View) scale(g *gocui.Gui) (int, int, int, int) {
	maxX, maxY := g.Size()
	return int(v.x0 * float32(maxX)),
		int(v.y0 * float32(maxY)),
		int(v.x1 * float32(maxX)),
		int(v.y1 * float32(maxY))
}

func (v *View) draw(g *gocui.Gui) (*gocui.View, error) {
	x0, y0, x1, y1 := v.scale(g)
	return g.SetView(v.name, x0, y0, x1, y1)
}

func (v *View) get(g *gocui.Gui) (*gocui.View, error) {
	return g.View(v.name)
}

func (v *View) active(g *gocui.Gui) error {
	_, err := g.SetCurrentView(v.name)
	if err != nil {
		return err
	}
	_, err = g.SetViewOnTop(v.name)
	if err != nil {
		return err
	}
	return nil
}
