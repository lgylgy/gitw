package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
	"github.com/lgylgy/gitw/gui"
)

type View interface {
	Draw(g *gocui.Gui) error
	Update(g *gocui.Gui, text string) error
	GetName() string
}

type Layout struct {
	views     []View
	onChanged chan string
}

func NewLayout(g *gocui.Gui, config *git.Config) *Layout {
	changed := make(chan string)
	layout := &Layout{
		views: []View{
			gui.NewSidebarView(g, changed, config),
			gui.NewBranchView(),
			gui.NewRemotesView(),
			gui.NewContentView(),
			gui.NewLogView(),
		},
		onChanged: changed,
	}
	go func() {
		for text := range changed {
			for _, view := range layout.views {
				view.Update(g, text)
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
