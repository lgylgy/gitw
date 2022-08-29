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

func CreateDefaultViews(g *gocui.Gui, repositories *git.Repositories, events chan<- *gui.Event) map[string]View {
	return map[string]View{
		"repositories": CreateView("repositories", g, repositories, events),
		"branch":       CreateView("branch", g, repositories, events),
		"remotes":      CreateView("remotes", g, repositories, events),
		"content":      CreateView("content", g, repositories, events),
	}
}

func CreateView(name string, g *gocui.Gui, repositories *git.Repositories, events chan<- *gui.Event) View {
	switch name {
	case "repositories":
		return gui.NewSidebarView(g, events, repositories)
	case "branch":
		return gui.NewBranchView()
	case "remotes":
		return gui.NewRemotesView()
	case "content":
		return gui.NewContentView()
	case "actions":
		return gui.NewActionsView(g, events)
	}
	return nil
}
