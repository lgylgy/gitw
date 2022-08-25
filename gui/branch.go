package gui

import (
	"fmt"

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
			y0:   0.74,
			x1:   0.5,
			y1:   0.89,
		},
	}
}

func (bv *BranchView) Draw(g *gocui.Gui) error {
	view, err := bv.View.draw(g)
	if err == gocui.ErrUnknownView {
		view.Title = "Current branch"
		return nil
	}
	return err
}

func (bv *BranchView) Update(g *gocui.Gui, text string) error {
	view, err := bv.View.get(g)
	if err != nil {
		return err
	}
	view.Clear()
	fmt.Fprint(view, text)
	return nil
}
