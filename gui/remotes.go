package gui

import (
	"fmt"

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
			y0:   0.74,
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

func (rv *RemotesView) Update(g *gocui.Gui, text string) error {
	view, err := rv.View.get(g)
	if err != nil {
		return err
	}
	view.Clear()
	fmt.Fprint(view, text)
	return nil
}
