package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
	"github.com/lgylgy/gitw/gui"
)

type Layout struct {
	g       *gocui.Gui
	manager *git.Manager
	views   map[string]View
	events  chan *gui.Event
	errors  chan error
}

func NewLayout(g *gocui.Gui, manager *git.Manager) *Layout {
	events := make(chan *gui.Event, 2)
	errors := make(chan error, 2)
	layout := &Layout{
		g:       g,
		manager: manager,
		views:   CreateDefaultViews(g, manager, events),
		events:  events,
		errors:  errors,
	}
	go layout.run()
	return layout
}

func (l *Layout) Manage() {
	for _, view := range l.views {
		err := view.Draw(l.g)
		if err != nil {
			l.errors <- fmt.Errorf("view %s: %v", view.GetName(), err)
		}
	}
}

func (l *Layout) run() {
	for {
		select {
		case event := <-l.events:
			switch event.T {
			case gui.Add:
				l.addView(event)
			case gui.Update:
				l.update(event)
			case gui.Remove:
				l.delete(event)
			}
		case err := <-l.errors:
			view := gui.NewErrorView(l.g, l.events, err)
			l.views[view.GetName()] = view
			l.update(&gui.Event{
				T:     gui.Update,
				Views: []string{view.GetName()},
			})
		}
	}
}

func (l *Layout) update(event *gui.Event) {
	if len(event.Views) == 0 {
		for _, view := range l.views {
			err := view.Update(l.g, event.Repo)
			if err != nil {
				l.errors <- fmt.Errorf("view %s: %v", view.GetName(), err)
			}
		}
		return
	}
	for _, name := range event.Views {
		view, ok := l.views[name]
		if ok {
			err := view.Update(l.g, event.Repo)
			if err != nil {
				l.errors <- fmt.Errorf("view %s: %v", view.GetName(), err)
			}
		}
	}
}

func (l *Layout) delete(event *gui.Event) {
	for _, name := range event.Views {
		view, ok := l.views[name]
		if ok {
			err := l.g.DeleteView(name)
			if err != nil {
				l.errors <- fmt.Errorf("view %s: %v", view.GetName(), err)
			}
			delete(l.views, name)
			l.update(&gui.Event{
				T:     gui.Update,
				Views: []string{"repositories"},
			})
		}
	}
}

func (l *Layout) addView(event *gui.Event) {
	for _, name := range event.Views {
		view := CreateView(name, l.g, l.manager, l.events)
		if view != nil {
			l.views[view.GetName()] = view
			err := view.Draw(l.g)
			if err != nil {
				l.errors <- fmt.Errorf("view %s: %v", view.GetName(), err)
			}
			err = view.Update(l.g, event.Repo)
			if err != nil {
				l.errors <- fmt.Errorf("view %s: %v", view.GetName(), err)
			}
		}
	}
}
