package gui

import (
	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
)

type LogView struct {
	View
}

func NewLogView() *LogView {
	return &LogView{
		View{
			name: "log",
			x0:   0,
			y0:   0.90,
			x1:   0.99,
			y1:   0.99,
		},
	}
}

func (lv *LogView) Draw(g *gocui.Gui) error {
	view, err := lv.View.draw(g)
	if err == gocui.ErrUnknownView {
		view.Title = "Log"
		return nil
	}
	return err
}

func (lv *LogView) Update(g *gocui.Gui, _ *git.Repository) error {
	_, err := lv.View.get(g)
	if err != nil {
		return err
	}
	return nil
}
