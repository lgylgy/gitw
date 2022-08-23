package gui

import (
	"log"

	"github.com/jroimartin/gocui"
)

type BranchView struct {
	Name string
	x0   int
	y0   int
	x1   int
	y1   int
}

func NewBranchView() *BranchView {
	return &BranchView{
		Name: "branch",
		x0:   0,
		y0:   21,
		x1:   50,
		y1:   26,
	}
}

func (sv *BranchView) Draw(g *gocui.Gui) {
	if v, err := g.SetView(sv.Name, sv.x0, sv.y0, sv.x1, sv.y1); err != nil {
		log.Println(v)
		log.Println(err)
		v.Title = "Current branch"
	}
}
