package gui

import (
	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
)

type ActionsView struct {
	View
	index  int
	events chan<- *Event
}

func NewActionsView(g *gocui.Gui, events chan<- *Event) *ActionsView {
	view := &ActionsView{
		View{
			name: "actions",
			x0:   0.3,
			y0:   0.3,
			x1:   0.7,
			y1:   0.7,
		},
		0,
		events,
	}
	return view
}

func (av *ActionsView) onChange(position int) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return nil
	}
}

func (av *ActionsView) Draw(g *gocui.Gui) error {
	view, err := av.View.draw(g)
	if err == gocui.ErrUnknownView {
		err := g.SetKeybinding(av.View.name, gocui.KeyArrowDown, gocui.ModNone, av.onChange(1))
		if err != nil {
			return err
		}
		err = g.SetKeybinding(av.View.name, gocui.KeyArrowUp, gocui.ModNone, av.onChange(-1))
		if err != nil {
			return err
		}
		err = g.SetKeybinding(av.View.name, gocui.KeyCtrlX, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			av.events <- &Event{
				T:    Remove,
				View: av.GetName(),
			}
			return nil
		})
		if err != nil {
			return err
		}
		err = av.active(g)
		if err != nil {
			return err
		}

		view.SelFgColor = gocui.ColorBlack
		view.SelBgColor = gocui.ColorGreen
		view.Highlight = true

		view.Title = "Actions"
		return nil
	}
	return err
}

func (av *ActionsView) Update(*gocui.Gui, *git.Repository) error {
	return nil
}
