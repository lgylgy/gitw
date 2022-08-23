package gui

import (
	"github.com/jroimartin/gocui"
)

type ContentView struct {
	Name string
	x0   int
	y0   int
	x1   int
	y1   int
}

func NewContentView() *ContentView {
	return &ContentView{
		Name: "main",
		x0:   21,
		y0:   0,
		x1:   100,
		y1:   20,
	}
}

func (cv *ContentView) Draw(g *gocui.Gui) {
	if _, err := g.SetView(cv.Name, cv.x0, cv.y0, cv.x1, cv.y1); err != nil {
	}
}
