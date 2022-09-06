package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
)

type SidebarView struct {
	View
	index   int
	manager *git.Manager
	events  chan<- *Event
}

func NewSidebarView(g *gocui.Gui, events chan<- *Event,
	manager *git.Manager) *SidebarView {
	view := &SidebarView{
		newView("repositories", 0, 0, 0.2, 0.7),
		0,
		manager,
		events,
	}
	return view
}

func (sbv *SidebarView) onChange(position int) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		newPosition := sbv.index + position
		if newPosition < 0 || newPosition >= len(sbv.manager.ListRepos()) {
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
		selected := sbv.manager.Select(sbv.index)
		sbv.events <- &Event{
			T:    Update,
			Repo: selected,
		}
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
		err = g.SetKeybinding(sbv.View.name, gocui.KeySpace, gocui.ModNone, func(*gocui.Gui, *gocui.View) error {
			sbv.events <- &Event{
				T:    Add,
				View: "actions",
			}
			return nil
		})
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

		for _, repo := range sbv.manager.ListRepos() {
			fmt.Fprintf(view, "%s\n", repo)
		}
		view.Title = "LGY repositories"
		return nil
	}
	return err
}

func (sbv *SidebarView) Update(g *gocui.Gui, _ *git.Repository) error {
	view, err := sbv.View.get(g)
	if err != nil {
		return err
	}
	_, err = g.SetCurrentView(sbv.View.name)
	if err != nil {
		return err
	}
	g.Update(func(g *gocui.Gui) error {
		view.Clear()
		for _, remote := range sbv.manager.ListRepos() {
			fmt.Fprintf(view, "%s\n", remote)
		}
		return nil
	})
	return err
}
