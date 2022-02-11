package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/gui"
)

var (
	sidebarView *gui.SidebarView
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()
	g.SetManagerFunc(layout)

	// Side bar
	sidebarView, err = gui.NewSidebarView()
	if err != nil {
		log.Fatalln(err)
	}

	// Quit binding
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Fatalln(err)
	}

	// Main loop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalln(err)
	}
}

func layout(g *gocui.Gui) error {
	sidebarView.Draw(g)
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
