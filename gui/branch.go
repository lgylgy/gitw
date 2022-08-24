package gui

import (
	"github.com/jroimartin/gocui"
)

type BranchView struct {
	View
}

func NewBranchView() *BranchView {
	return &BranchView{
		View{
			name: "branch",
			x0:   0,
			y0:   0.71,
			x1:   0.5,
			y1:   0.89,
		},
	}
}

func (sv *BranchView) Draw(g *gocui.Gui) error {
	view, err := sv.View.draw(g)
	if err == gocui.ErrUnknownView {
		view.Title = "Current branch"
		return nil
	}
	return err
}
