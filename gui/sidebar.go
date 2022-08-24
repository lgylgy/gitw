package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type SidebarView struct {
	View
}

func NewSidebarView() *SidebarView {
	return &SidebarView{
		View{
			name: "repositories",
			x0:   0,
			y0:   0,
			x1:   20,
			y1:   20,
		},
	}
}

func (sbv *SidebarView) Draw(g *gocui.Gui) error {
	view, err := sbv.View.Draw(g)
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
