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
	views     []View
	onChanged chan *git.Repository
}

func NewLayout(g *gocui.Gui, repositories *git.Repositories) *Layout {
	changed := make(chan *git.Repository)
	layout := &Layout{
		views: []View{
			gui.NewSidebarView(g, changed, repositories),
			gui.NewBranchView(),
			gui.NewRemotesView(),
			gui.NewContentView(),
			gui.NewLogView(),
		},
		onChanged: changed,
	}
	go func() {
		for repository := range changed {
			for _, view := range layout.views {
				_ = view.Update(g, repository)
			}
		}
	}()
	return layout
}

func (l *Layout) Draw(g *gocui.Gui) error {
	for _, view := range l.views {
		err := view.Draw(g)
		if err != nil {
			return fmt.Errorf("view %s: %v", view.GetName(), err)
		}
	}
	return nil
}
