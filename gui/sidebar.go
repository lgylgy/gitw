package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type SidebarView struct {
	Name string
	x0   int
	y0   int
	x1   int
	y1   int
}

func NewSidebarView() (*SidebarView, error) {
	return &SidebarView{
		Name: "repositories",
		x0:   0,
		y0:   0,
		x1:   20,
		y1:   20,
	}, nil
}

func (sB *SidebarView) Draw(g *gocui.Gui) {
	if v, err := g.SetView(sB.Name, sB.x0, sB.y0, sB.x1, sB.y1); err != nil {
		v.SelFgColor = gocui.ColorBlack
		v.SelBgColor = gocui.ColorGreen
		v.Highlight = true

		repositories := []string{"repo1", "repo2", "repo3"}
		for _, item := range repositories {
			fmt.Fprintf(v, "%s\n", item)
		}
		v.Title = "LGY repositories"
	}
}
