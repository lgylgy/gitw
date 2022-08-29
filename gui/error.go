package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
)

type ErrorView struct {
	View
	events chan<- *Event
	err    error
}

func NewErrorView(g *gocui.Gui, events chan<- *Event, err error) *ErrorView {
	return &ErrorView{
		View{
			name: "log",
			x0:   0.1,
			y0:   0.1,
			x1:   0.9,
			y1:   0.9,
		},
		events,
		err,
	}
}

func (ev *ErrorView) Draw(g *gocui.Gui) error {
	view, err := ev.View.draw(g)
	if err == gocui.ErrUnknownView {
		view.Title = "Log"
		view.Wrap = true
		_, err = g.SetCurrentView(ev.View.GetName())
		if err != nil {
			return err
		}
		_, err = g.SetViewOnTop(ev.View.GetName())
		if err != nil {
			return err
		}
		err = g.SetKeybinding(ev.View.name, gocui.KeyCtrlX, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			ev.events <- &Event{
				T:    Remove,
				View: ev.GetName(),
			}
			return nil
		})
		return err
	}
	return err
}

func (ev *ErrorView) Update(g *gocui.Gui, _ *git.Repository) error {
	view, err := ev.View.get(g)
	if err != nil {
		return err
	}
	g.Update(func(g *gocui.Gui) error {
		view.Clear()
		fmt.Fprintf(view, "%s\n", ev.err)
		return nil
	})
	return nil
}