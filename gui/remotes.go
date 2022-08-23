package gui

import (
	"github.com/jroimartin/gocui"
)

type RemotesView struct {
	Name string
	x0   int
	y0   int
	x1   int
	y1   int
}

func NewRemotesView() *RemotesView {
	return &RemotesView{
		Name: "remotes",
		x0:   51,
		y0:   21,
		x1:   100,
		y1:   26,
	}
}

func (rv *RemotesView) Draw(g *gocui.Gui) {
	if v, err := g.SetView(rv.Name, rv.x0, rv.y0, rv.x1, rv.y1); err != nil {
		v.Title = "Remotes"
	}
}
