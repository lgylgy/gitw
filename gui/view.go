package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func newView(name string, x0, y0, x1, y1 float32) View {
	return View{
		events: make(chan *event, 5),
		id:     0,
		name:   name,
		x0:     x0,
		y0:     y0,
		x1:     x1,
		y1:     y1,
	}
}

type event struct {
	id     int32
	err    error
	output string
}

type View struct {
	events chan *event
	id     int32
	name   string
	x0     float32
	y0     float32
	x1     float32
	y1     float32
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

func (v *View) update(fun func() (string, error)) {
	v.id++
	go func(id int32) {
		output, err := fun()
		v.events <- &event{
			id:     id,
			err:    err,
			output: output,
		}
	}(v.id)
}

func (v *View) process(g *gocui.Gui) {
	for event := range v.events {
		if event.id == v.id {
			view, err := g.View(v.name)
			if err != nil {
				return
			}
			g.Update(func(g *gocui.Gui) error {
				view.Clear()
				fmt.Fprintf(view, "%s\n", event.output)
				return nil
			})
		}
	}
}
