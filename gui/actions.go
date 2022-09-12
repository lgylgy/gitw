package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
)

type ActionsView struct {
	View
	index   int
	manager *git.Manager
	repo    *git.Repository

	events  chan<- *Event
	results chan *git.Result
}

func NewActionsView(events chan<- *Event, manager *git.Manager, results chan *git.Result) *ActionsView {
	return &ActionsView{
		newView("actions", 0.2, 0.1, 0.8, 0.4),
		0,
		manager,
		nil,
		events,
		results,
	}
}

func (av *ActionsView) onChange(position int) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		newPosition := av.index + position
		if newPosition < 0 || newPosition >= len(av.manager.ListActions()) {
			return nil
		}
		view, err := g.View(av.View.name)
		if err != nil {
			return err
		}
		err = view.SetCursor(0, newPosition)
		if err != nil {
			return err
		}
		av.index = newPosition
		return nil
	}
}

func (av *ActionsView) onProcess() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if av.repo == nil {
			return nil
		}

		action := av.manager.GetAction(av.index)
		if action != nil {
			result := make(chan bool)
			go func() {
				defer func() {
					close(result)
				}()
				action.Process(av.repo.Path, av.results)
			}()
			<-result
		}
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
		err = g.SetKeybinding(av.View.name, gocui.KeyEnter, gocui.ModNone, av.onProcess())
		if err != nil {
			return err
		}
		err = g.SetKeybinding(av.View.name, gocui.KeySpace, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			av.events <- &Event{
				T:     Remove,
				Views: []string{av.GetName(), "output"},
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

func (av *ActionsView) Update(g *gocui.Gui, repo *git.Repository) error {
	av.repo = repo
	view, err := av.View.get(g)
	if err != nil {
		return err
	}
	g.Update(func(g *gocui.Gui) error {
		view.Clear()
		for _, repo := range av.manager.ListActions() {
			fmt.Fprintf(view, "* %s\n", repo)
		}
		return nil
	})
	return nil
}
