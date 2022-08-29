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
	l.events <- &gui.Event{
		T: gui.Draw,
	}
}

func (l *Layout) run() {
	for {
		select {
		case event := <-l.events:
			switch event.T {
			case gui.Add:
				l.addView(event)
			case gui.Draw:
				l.draw(event)
			case gui.Update:
				l.update(event)
			case gui.Remove:
				l.delete(event)
			}
		case err := <-l.errors:
			view := gui.NewErrorView(l.g, l.events, err)
			l.views[view.GetName()] = view
			l.events <- &gui.Event{
				T:    gui.Update,
				View: view.GetName(),
			}
		}
	}
}

func (l *Layout) draw(event *gui.Event) {
	if event.View == "" {
		for _, view := range l.views {
			err := view.Draw(l.g)
			if err != nil {
				l.errors <- fmt.Errorf("view %s: %v", view.GetName(), err)
			}
		}
		return
	}
	view, ok := l.views[event.View]
	if ok {
		err := view.Draw(l.g)
		if err != nil {
			l.errors <- fmt.Errorf("view %s: %v", view.GetName(), err)
		}
	}
}

func (l *Layout) update(event *gui.Event) {
	if event.View == "" {
		for _, view := range l.views {
			err := view.Update(l.g, event.Repo)
			if err != nil {
				l.errors <- fmt.Errorf("view %s: %v", view.GetName(), err)
			}
		}
		return
	}
	view, ok := l.views[event.View]
	if ok {
		err := view.Update(l.g, event.Repo)
		if err != nil {
			l.errors <- fmt.Errorf("view %s: %v", view.GetName(), err)
		}
	}
}

func (l *Layout) delete(event *gui.Event) {
	view, ok := l.views[event.View]
	if ok {
		err := l.g.DeleteView(event.View)
		if err != nil {
			l.errors <- fmt.Errorf("view %s: %v", view.GetName(), err)
		}
		delete(l.views, event.View)
		l.events <- &gui.Event{
			T:    gui.Update,
			View: "repositories",
		}
	}
}

func (l *Layout) addView(event *gui.Event) {
	view := CreateView(event.View, l.g, l.manager, l.events)
	if view != nil {
		l.views[view.GetName()] = view
		l.events <- &gui.Event{
			T: gui.Draw,
		}
		l.events <- &gui.Event{
			T:    gui.Update,
			View: view.GetName(),
		}
		return
	}
	l.errors <- fmt.Errorf("unknown view %s", event.View)
}
