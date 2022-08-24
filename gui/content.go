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
			x0:   0.21,
			y0:   0,
			x1:   0.99,
			y1:   0.7,
		},
	}
}

func (cv *ContentView) Draw(g *gocui.Gui) error {
	_, err := cv.View.draw(g)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	return nil
}
