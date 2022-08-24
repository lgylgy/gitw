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

func NewSidebarView() *SidebarView {
	return &SidebarView{
		Name: "repositories",
		x0:   0,
		y0:   0,
		x1:   20,
		y1:   20,
	}
}

func (sbv *SidebarView) Draw(g *gocui.Gui) error {
	view, err := g.SetView(sbv.Name, sbv.x0, sbv.y0, sbv.x1, sbv.y1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		view.SelFgColor = gocui.ColorBlack
		view.SelBgColor = gocui.ColorGreen
		view.Highlight = true

		repositories := []string{"repo1", "repo2", "repo3"}
		for _, item := range repositories {
			fmt.Fprintf(view, "%s\n", item)
		}
		view.Title = "LGY repositories"
	}
	return nil
}
