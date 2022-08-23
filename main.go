package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/gui"
)

var (
	sidebarView *gui.SidebarView
	branchView  *gui.BranchView
	remotesView *gui.RemotesView
	contentView *gui.ContentView
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()
	g.SetManagerFunc(layout)

	// Side bar
	sidebarView = gui.NewSidebarView()
	if err != nil {
		log.Fatalln(err)
	}

	// Branch layout
	branchView = gui.NewBranchView()
	if err != nil {
		log.Fatalln(err)
	}

	// Remote layout
	remotesView = gui.NewRemotesView()
	if err != nil {
		log.Fatalln(err)
	}

	// Main content
	contentView = gui.NewContentView()
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
	branchView.Draw(g)
	remotesView.Draw(g)
	contentView.Draw(g)
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
