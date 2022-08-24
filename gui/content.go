package gui

import (
	"github.com/jroimartin/gocui"
)

type ContentView struct {
	View
}

func NewContentView() *ContentView {
	return &ContentView{
		View{
			name: "main",
			x0:   21,
			y0:   0,
			x1:   100,
			y1:   20,
		},
	}
}

func (cv *ContentView) Draw(g *gocui.Gui) error {
	_, err := cv.View.Draw(g)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	return nil
}
