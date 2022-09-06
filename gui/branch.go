package gui

import (
	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
)

type BranchView struct {
	View
}

func NewBranchView(g *gocui.Gui) *BranchView {
	view := &BranchView{
		newView("branch", 0, 0.74, 0.6, 0.99),
	}
	go view.process(g)
	return view
}

func (bv *BranchView) Draw(g *gocui.Gui) error {
	view, err := bv.View.draw(g)
	if err == gocui.ErrUnknownView {
		view.Title = "Current branch"
		view.Wrap = true
		return nil
	}
	return err
}

func (bv *BranchView) Update(g *gocui.Gui, current *git.Repository) error {
	view, err := bv.View.get(g)
	if err != nil {
		return err
	}
	view.Clear()
	bv.update(func() (string, error) {
		name, err := current.GetCurrentBranch()
		if err != nil {
			return "", err
		}
		return name, nil
	})
	return nil
}
