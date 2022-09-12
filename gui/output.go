package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
)

type OutputView struct {
	View
	results chan *git.Result
}

func NewOutputView(g *gocui.Gui, results chan *git.Result) *OutputView {
	view := &OutputView{
		newView("output", 0.2, 0.5, 0.8, 0.8),
		results,
	}
	go func() {
		for result := range results {
			view, err := view.View.get(g)
			if err != nil {
				continue
			}
			fmt.Fprintf(view, "%s\n", result.Output)
			if result.Err != nil {
				fmt.Fprintf(view, "%s\n", result.Err.Error())
			}
		}
	}()
	return view
}

func (ov *OutputView) Draw(g *gocui.Gui) error {
	view, err := ov.View.draw(g)
	if err == gocui.ErrUnknownView {
		view.SelFgColor = gocui.ColorBlack
		view.SelBgColor = gocui.ColorGreen

		view.Title = "Output"
		return nil
	}
	return err
}

func (ov *OutputView) Update(g *gocui.Gui, _ *git.Repository) error {
	view, err := ov.View.get(g)
	if err != nil {
		return err
	}
	view.Clear()

	return nil
}
