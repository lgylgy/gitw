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
			x0:   51,
			y0:   21,
			x1:   100,
			y1:   26,
		},
	}
}

func (rv *RemotesView) Draw(g *gocui.Gui) error {
	view, err := rv.View.Draw(g)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		view.Title = "Remotes"
	}
	return nil
}
