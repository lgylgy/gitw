package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
)

type RemotesView struct {
	View
}

func NewRemotesView() *RemotesView {
	return &RemotesView{
		View{
			name: "remotes",
			x0:   0.61,
			y0:   0.74,
			x1:   0.99,
			y1:   0.99,
		},
	}
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
	remotes, err := current.GetRemotes()
	if err != nil {
		return err
	}
	g.Update(func(g *gocui.Gui) error {
		view.Clear()
		for _, remote := range remotes {
			fmt.Fprintf(view, "%s\n", remote)
		}
		return nil
	})
	return nil
}
