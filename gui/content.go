package gui

import (
	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
)

type ContentView struct {
	View
}

func NewContentView(g *gocui.Gui) *ContentView {
	view := &ContentView{
		newView("content", 0.21, 0, 0.99, 0.7),
	}
	go view.process(g)
	return view
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
	cv.update(func() (string, error) {
		commits, err := current.GetCommits()
		if err != nil {
			return "", err
		}
		return commits, nil
	})
	return nil
}
