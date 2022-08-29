package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
)

type ContentView struct {
	View
}

func NewContentView() *ContentView {
	return &ContentView{
		View{
			name: "content",
			x0:   0.21,
			y0:   0,
			x1:   0.99,
			y1:   0.7,
		},
	}
}

func (cv *ContentView) Draw(g *gocui.Gui) error {
	view, err := cv.View.draw(g)
	if err == gocui.ErrUnknownView {
		view.Wrap = true
		return nil
	}
	return err
}

func (cv *ContentView) Update(g *gocui.Gui, current *git.Repository) error {
	view, err := cv.View.get(g)
	if err != nil {
		return err
	}
	view.Clear()
	commits, err := current.GetCommits()
	if err != nil {
		return err
	}
	g.Update(func(g *gocui.Gui) error {
		view.Clear()
		fmt.Fprintf(view, "%s\n", commits)
		return nil
	})
	return nil
}
