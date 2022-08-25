package gui

import (
	"fmt"

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
	if err == gocui.ErrUnknownView {
		return nil
	}
	return err
}

func (cv *ContentView) Update(g *gocui.Gui, text string) error {
	view, err := cv.View.get(g)
	if err != nil {
		return err
	}
	view.Clear()
	fmt.Fprint(view, text)
	return nil
}
