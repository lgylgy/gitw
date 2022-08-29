package main

import (
	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
	"github.com/lgylgy/gitw/gui"
)

type View interface {
	Draw(*gocui.Gui) error
	Update(*gocui.Gui, *git.Repository) error
	GetName() string
}

func CreateDefaultViews(g *gocui.Gui, manager *git.Manager, events chan<- *gui.Event) map[string]View {
	return map[string]View{
		"repositories": CreateView("repositories", g, manager, events),
		"branch":       CreateView("branch", g, manager, events),
		"remotes":      CreateView("remotes", g, manager, events),
		"content":      CreateView("content", g, manager, events),
	}
}

func CreateView(name string, g *gocui.Gui, manager *git.Manager, events chan<- *gui.Event) View {
	switch name {
	case "repositories":
		return gui.NewSidebarView(g, events, manager)
	case "branch":
		return gui.NewBranchView()
	case "remotes":
		return gui.NewRemotesView()
	case "content":
		return gui.NewContentView()
	case "actions":
		return gui.NewActionsView(g, events, manager)
	}
	return nil
}
