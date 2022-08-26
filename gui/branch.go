package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
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
			x1:   0.6,
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

func (bv *BranchView) Update(g *gocui.Gui, current *git.Repository) error {
	view, err := bv.View.get(g)
	if err != nil {
		return err
	}
	name, err := current.GetCurrentBranch()
	if err != nil {
		return err
	}
	g.Update(func(g *gocui.Gui) error {
		view.Clear()
		fmt.Fprintf(view, "%s\n", name)
		return nil
	})
	return nil
}
