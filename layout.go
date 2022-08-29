package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
	"github.com/lgylgy/gitw/gui"
)

type View interface {
	Draw(*gocui.Gui) error
	Update(*gocui.Gui, *git.Repository) error
	GetName() string
}

type Layout struct {
	g      *gocui.Gui
	views  map[string]View
	events chan *gui.Event
	errors chan error
}

func NewLayout(g *gocui.Gui, repositories *git.Repositories) *Layout {
	events := make(chan *gui.Event, 2)
	errors := make(chan error, 2)
	layout := &Layout{
		g: g,
		views: map[string]View{
			"repositories": gui.NewSidebarView(g, events, repositories),
			"branch":       gui.NewBranchView(),
			"remotes":      gui.NewRemotesView(),
			"content":      gui.NewContentView(),
		},
		events: events,
		errors: errors,
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
			case gui.Draw, gui.Add:
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
				T:    gui.Draw,
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
