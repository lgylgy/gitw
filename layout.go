package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/gui"
)

type View interface {
	Draw(g *gocui.Gui) error
	GetName() string
}

type Layout struct {
	views []View
}

func NewLayout(g *gocui.Gui) *Layout {
	return &Layout{
		views: []View{
			gui.NewSidebarView(g),
			gui.NewBranchView(),
			gui.NewRemotesView(),
			gui.NewContentView(),
			gui.NewLogView(),
		},
	}
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
