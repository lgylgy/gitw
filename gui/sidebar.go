package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type SidebarView struct {
	View
	current int
	repos   []string
}

func NewSidebarView(g *gocui.Gui) *SidebarView {
	view := &SidebarView{
		View{
			name: "repositories",
			x0:   0,
			y0:   0,
			x1:   0.2,
			y1:   0.7,
		},
		0,
		[]string{"repo1", "repo2", "repo3"},
	}
	return view
}

func (sbv *SidebarView) onChange(position int) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		newPosition := sbv.current + position
		if newPosition < 0 || newPosition >= len(sbv.repos) {
			return nil
		}
		view, err := g.View(sbv.View.name)
		if err != nil {
			return err
		}
		err = view.SetCursor(0, newPosition)
		if err != nil {
			return err
		}
		sbv.current = newPosition
		return nil
	}
}

func (sbv *SidebarView) Draw(g *gocui.Gui) error {
	view, err := sbv.View.draw(g)
	if err == gocui.ErrUnknownView {

		err := g.SetKeybinding(sbv.View.name, gocui.KeyArrowDown, gocui.ModNone, sbv.onChange(1))
		if err != nil {
			return err
		}
		err = g.SetKeybinding(sbv.View.name, gocui.KeyArrowUp, gocui.ModNone, sbv.onChange(-1))
		if err != nil {
			return err
		}

		_, err = g.SetCurrentView(sbv.View.name)
		if err != nil {
			return err
		}

		view.SelFgColor = gocui.ColorBlack
		view.SelBgColor = gocui.ColorGreen
		view.Highlight = true
		for _, item := range sbv.repos {
			fmt.Fprintf(view, "%s\n", item)
		}
		view.Title = "LGY repositories"
		return nil
	}
	return err
}
