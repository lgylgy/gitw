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
			y0:   21,
			x1:   50,
			y1:   26,
		},
	}
}

func (sv *BranchView) Draw(g *gocui.Gui) error {
	view, err := sv.View.Draw(g)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		view.Title = "Current branch"
	}
	return nil
}
