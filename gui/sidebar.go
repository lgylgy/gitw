package gui

import (
	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
)

type SidebarView struct {
	View
	index        int
	repositories *git.Repositories
	events       chan<- *Event
}

func NewSidebarView(g *gocui.Gui, events chan<- *Event,
	repositories *git.Repositories) *SidebarView {
	view := &SidebarView{
		View{
			name: "repositories",
			x0:   0,
			y0:   0,
			x1:   0.2,
			y1:   0.7,
		},
		0,
		repositories,
		events,
	}
	return view
}

func (sbv *SidebarView) onChange(position int) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		newPosition := sbv.index + position
		if newPosition < 0 || newPosition >= sbv.repositories.Count() {
			return nil
		}
		view, err := g.View(sbv.View.name)
		if err != nil {
			return err
		}
		err = view.SetCursor(0, newPosition)
		if err != nil {
			return err
		}
		sbv.index = newPosition
		sbv.events <- &Event{
			T:    Update,
			Repo: sbv.repositories.Get(sbv.index),
		}
		sbv.repositories.Get(sbv.index)
		return nil
	}
}

func (sbv *SidebarView) Draw(g *gocui.Gui) error {
	view, err := sbv.View.draw(g)
	if err == gocui.ErrUnknownView {

		err := g.SetKeybinding(sbv.View.name, gocui.KeyArrowDown, gocui.ModNone, sbv.onChange(1))
		if err != nil {
			return err
		}
		err = g.SetKeybinding(sbv.View.name, gocui.KeyArrowUp, gocui.ModNone, sbv.onChange(-1))
		if err != nil {
			return err
		}
		_, err = g.SetCurrentView(sbv.View.name)
		if err != nil {
			return err
		}

		view.SelFgColor = gocui.ColorBlack
		view.SelBgColor = gocui.ColorGreen
		view.Highlight = true

		sbv.repositories.Display(view)
		view.Title = "LGY repositories"
		return nil
	}
	return err
}

func (sbv *SidebarView) Update(g *gocui.Gui, _ *git.Repository) error {
	_, err := g.SetCurrentView(sbv.View.name)
	return err
}
