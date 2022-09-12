package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
)

type DiffView struct {
	View
	cached  bool
	manager *git.Manager
	events  chan<- *Event
}

func getLabel(cached bool) string {
	if cached {
		return "Staged Changes"
	}
	return "Unstaged Changes"
}

func getName(cached bool) string {
	if cached {
		return "staged-changes"
	}
	return "unstaged-changes"
}

func getPositions(cached bool) (float32, float32, float32, float32) {
	if cached {
		return 0.51, 0.05, 0.95, 0.9
	}
	return 0.05, 0.05, 0.5, 0.9
}

func NewDiffView(events chan<- *Event, manager *git.Manager, cached bool) *DiffView {
	x0, y0, x1, y1 := getPositions(cached)
	return &DiffView{
		newView(getName(cached), x0, y0, x1, y1),
		cached,
		manager,
		events,
	}
}

func (dv *DiffView) Draw(g *gocui.Gui) error {
	view, err := dv.View.draw(g)
	if err == gocui.ErrUnknownView {
		err = g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			dv.events <- &Event{
				T:     Remove,
				Views: []string{getName(false), getName(true)},
			}
			return nil
		})
		if err != nil {
			return err
		}
		err = dv.active(g)
		if err != nil {
			return err
		}

		view.SelFgColor = gocui.ColorBlack
		view.SelBgColor = gocui.ColorGreen

		view.Title = getLabel(dv.cached)
		return nil
	}
	return err
}

func (dv *DiffView) Update(g *gocui.Gui, current *git.Repository) error {
	view, err := dv.View.get(g)
	if err != nil {
		return err
	}
	view.Clear()
	commits, err := current.GetDiff(dv.cached)
	if err != nil {
		return err
	}
	g.Update(func(g *gocui.Gui) error {
		view.Clear()
		fmt.Fprintf(view, "%s\n", commits)
		return nil
	})
	return nil
}
