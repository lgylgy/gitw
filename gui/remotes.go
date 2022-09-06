package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
)

type RemotesView struct {
	View
}

func NewRemotesView(g *gocui.Gui) *RemotesView {
	view := &RemotesView{
		newView("remotes", 0.61, 0.74, 0.99, 0.99),
	}
	go view.process(g)
	return view
}

func (rv *RemotesView) Draw(g *gocui.Gui) error {
	view, err := rv.View.draw(g)
	if err == gocui.ErrUnknownView {
		view.Title = "Remotes"
		view.Wrap = true
		return nil
	}
	return err
}

func (rv *RemotesView) Update(g *gocui.Gui, current *git.Repository) error {
	view, err := rv.View.get(g)
	if err != nil {
		return err
	}
	view.Clear()
	rv.update(func() (string, error) {
		remotes, err := current.GetRemotes()
		if err != nil {
			return "", err
		}
		result := ""
		for _, remote := range remotes {
			result += fmt.Sprintf("%s\n", remote)
		}
		return result, nil
	})
	return nil
}
