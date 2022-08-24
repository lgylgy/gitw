package gui

import (
	"github.com/jroimartin/gocui"
)

type RemotesView struct {
	View
}

func NewRemotesView() *RemotesView {
	return &RemotesView{
		View{
			name: "remotes",
			x0:   0.51,
			y0:   0.71,
			x1:   0.99,
			y1:   0.89,
		},
	}
}

func (rv *RemotesView) Draw(g *gocui.Gui) error {
	view, err := rv.View.draw(g)
	if err == gocui.ErrUnknownView {
		view.Title = "Remotes"
		return nil
	}
	return err
}
